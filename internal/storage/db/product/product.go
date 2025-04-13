package product

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"pvz-service/internal/entities"
	entitymap "pvz-service/internal/mappers/entity"
	"pvz-service/internal/models"
	"pvz-service/internal/storage/db/dberrors"
	"pvz-service/internal/storage/db/shared/consts"
	"pvz-service/internal/storage/db/shared/helper"

	sq "github.com/Masterminds/squirrel"
)

type ProductDB interface {
	AddProduct(ctx context.Context, addProduct *entities.AddProduct) (*models.Product, error)
	DeleteLastProduct(ctx context.Context, deleteProduct *entities.DeleteProduct) error
}

type productStorage struct {
	db     *sql.DB
	helper helper.DBHepler
	logger *slog.Logger
}

func New(db *sql.DB, logger *slog.Logger) ProductDB {
	return &productStorage{
		db:     db,
		helper: helper.New(db),
		logger: logger,
	}
}

func (s *productStorage) AddProduct(ctx context.Context, addProduct *entities.AddProduct) (*models.Product, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	isOpenedReception, receptionID, err := s.helper.CheckOpenedReceptions(ctx, tx, addProduct.PvzID)
	if err != nil {
		txErr := tx.Rollback() // For safety.
		if txErr != nil {
			s.logger.Error("tx rollback error: " + txErr.Error())
		}
		return nil, err
	}

	var product *entities.Product

	if isOpenedReception {
		product, err = s.insertProduct(ctx, tx, receptionID, addProduct.Type)
		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				s.logger.Error("tx rollback error: " + txErr.Error())
			}
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	if !isOpenedReception {
		return nil, dberrors.ErrInsertImpossible
	}

	return entitymap.MapToProduct(product), nil
}

func (s *productStorage) DeleteLastProduct(ctx context.Context, deleteProduct *entities.DeleteProduct) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	isOpenedReception, receptionID, err := s.helper.CheckOpenedReceptions(ctx, tx, deleteProduct.PvzID)
	if err != nil {
		txErr := tx.Rollback() // For safety.
		if txErr != nil {
			s.logger.Error("tx rollback error: " + txErr.Error())
		}
		return err
	}

	var affected int64
	if isOpenedReception {
		affected, err = s.deleteLastProduct(ctx, tx, receptionID)
		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				s.logger.Error("tx rollback error: " + txErr.Error())
			}
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	if !isOpenedReception {
		return dberrors.ErrDeleteImpossible
	}

	if affected == 0 {
		return dberrors.ErrNothingToDelete
	}

	return nil
}

func (s *productStorage) deleteLastProduct(ctx context.Context, tx *sql.Tx, receptionID int) (int64, error) {
	subq := fmt.Sprintf("(SELECT %s FROM %s WHERE %s=%d ORDER BY %s DESC LIMIT 1)",
		consts.ID, consts.ProductsTable, consts.ProductReceptionID, receptionID, consts.ProductDateTime)

	query, args, err := sq.Delete(consts.ProductsTable).
		Where(fmt.Sprintf("%s IN %s", consts.ID, subq)).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return 0, err
	}

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affected, nil
}

func (s *productStorage) insertProduct(ctx context.Context, tx *sql.Tx, receptionID int, pType string) (*entities.Product, error) {
	query, args, err := sq.Insert(consts.ProductsTable).
		Columns(consts.ProductReceptionID, consts.ProductType).
		Values(receptionID, pType).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var product entities.Product

	row := tx.QueryRowContext(ctx, query, args...)
	err = row.Scan(&product.ID, &product.DateTime, &product.ReceptionID, &product.Type)
	if err != nil {
		return nil, s.helper.CatchPQErrors(err)
	}

	return &product, nil
}

package product

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"pvz-service/internal/entities"
	"pvz-service/internal/enum"
	entitymap "pvz-service/internal/mappers/entity"
	"pvz-service/internal/models"
	"pvz-service/internal/storage/db/dberrors"
	"pvz-service/internal/storage/db/shared/consts"

	sq "github.com/Masterminds/squirrel"

	"github.com/lib/pq"
)

type ProductDB interface {
	AddProduct(ctx context.Context, addProduct *entities.AddProduct) (*models.Product, error)
	DeleteLastProduct(ctx context.Context, deleteProduct *entities.DeleteProduct) error
}

type productStorage struct {
	db     *sql.DB
	logger *slog.Logger
}

func New(db *sql.DB, logger *slog.Logger) ProductDB {
	return &productStorage{
		db:     db,
		logger: logger,
	}
}

func (s *productStorage) AddProduct(ctx context.Context, addProduct *entities.AddProduct) (*models.Product, error) {
	selectQuery, selArgs, err := sq.Select(consts.ID).
		From(consts.ReceptionsTable).
		Where(sq.Eq{
			consts.ReceptionPvzID:  addProduct.PvzID,
			consts.ReceptionStatus: enum.StatusInProgress.String(),
		}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	var receptionID int
	row := tx.QueryRowContext(ctx, selectQuery, selArgs...)
	selectErr := row.Scan(&receptionID)

	var product entities.Product

	if selectErr == nil {
		insertQuery, insArgs, err := sq.Insert(consts.ProductsTable).
			Columns(consts.ProductReceptionID, consts.ProductType).
			Values(receptionID, addProduct.Type).
			Suffix("RETURNING *").
			PlaceholderFormat(sq.Dollar).ToSql()
		if err != nil {
			return nil, err
		}

		row = tx.QueryRowContext(ctx, insertQuery, insArgs...)
		err = row.Scan(&product.ID, &product.DateTime, &product.ReceptionID, &product.Type)
		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				s.logger.Error("tx rollback error: " + txErr.Error())
			}
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code.Name() == consts.PQInvalidTextRepresentation {
					return nil, errors.Join(dberrors.ErrEnumTypeViolation, err) // For mode specifics in logs.
				}
			}
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	if errors.Is(selectErr, sql.ErrNoRows) {
		return nil, dberrors.ErrInsertImpossible
	} else if selectErr != nil {
		return nil, err
	}

	return entitymap.MapToProduct(&product), nil
}

func (s *productStorage) DeleteLastProduct(ctx context.Context, deleteProduct *entities.DeleteProduct) error {
	selectQuery, selArgs, err := sq.Select(consts.ID).
		From(consts.ReceptionsTable).
		Where(sq.Eq{
			consts.ReceptionPvzID:  deleteProduct.PvzID,
			consts.ReceptionStatus: enum.StatusInProgress.String(),
		}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var receptionID int

	row := tx.QueryRowContext(ctx, selectQuery, selArgs...)
	selectErr := row.Scan(&receptionID)
	if selectErr == nil {
		subq := fmt.Sprintf("(SELECT %s FROM %s WHERE %s=%d ORDER BY %s DESC LIMIT 1)",
			consts.ID, consts.ProductsTable, consts.ProductReceptionID, receptionID, consts.ProductDateTime)

		deleteQuery, delArgs, err := sq.Delete(consts.ProductsTable).
			Where(fmt.Sprintf("%s IN %s", consts.ID, subq)).
			PlaceholderFormat(sq.Dollar).ToSql()
		if err != nil {
			return err
		}

		delRes, err := tx.ExecContext(ctx, deleteQuery, delArgs...)
		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				s.logger.Error("tx rollback error: " + txErr.Error())
			}
			return err
		}

		err = tx.Commit()
		if err != nil {
			return err
		}

		affected, err := delRes.RowsAffected()
		if err != nil {
			return err
		}

		if affected == 0 {
			return dberrors.ErrNothingToDelete
		}

		return nil
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	if selectErr == sql.ErrNoRows {
		return dberrors.ErrDeleteImpossible
	}

	return selectErr
}

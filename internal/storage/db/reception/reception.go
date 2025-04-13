package reception

import (
	"context"
	"database/sql"
	"log/slog"
	"pvz-service/internal/entities"
	"pvz-service/internal/enum"
	"pvz-service/internal/models"
	"pvz-service/internal/storage/db/dberrors"
	"pvz-service/internal/storage/db/shared/consts"
	"pvz-service/internal/storage/db/shared/helper"

	entitymap "pvz-service/internal/mappers/entity"

	sq "github.com/Masterminds/squirrel"
)

type ReceptionDB interface {
	CreateReception(ctx context.Context, createreception *entities.CreateReception) (*models.Reception, error)
	CloseReception(ctx context.Context, closeReception *entities.CloseReception) (*models.Reception, error)
}

type receptionStorage struct {
	db     *sql.DB
	helper helper.DBHepler
	logger *slog.Logger
}

func New(db *sql.DB, logger *slog.Logger) ReceptionDB {
	return &receptionStorage{
		db:     db,
		helper: helper.New(db),
		logger: logger,
	}
}

func (s *receptionStorage) CreateReception(ctx context.Context, createReception *entities.CreateReception) (*models.Reception, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	isOpenedReception, _, err := s.helper.CheckOpenedReceptions(ctx, tx, createReception.PvzID)
	if err != nil {
		txErr := tx.Rollback() // For safety.
		if txErr != nil {
			s.logger.Error("tx rollback error: " + txErr.Error())
		}
		return nil, err
	}

	var reception *entities.Reception

	if !isOpenedReception {
		reception, err = s.insertNewReception(ctx, tx, createReception.PvzID)
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

	if isOpenedReception {
		return nil, dberrors.ErrInsertImpossible
	}

	return entitymap.MapToReception(reception, nil), nil
}

func (s *receptionStorage) CloseReception(ctx context.Context, closeReception *entities.CloseReception) (*models.Reception, error) {
	query, args, err := sq.Update(consts.ReceptionsTable).
		Set(consts.ReceptionStatus, enum.StatusClose.String()).
		Where(sq.Eq{
			consts.ReceptionPvzID:  closeReception.PvzID,
			consts.ReceptionStatus: enum.StatusInProgress.String(),
		}).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var reception entities.Reception

	row := s.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(&reception.ID, &reception.DateTime, &reception.PvzID, &reception.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dberrors.ErrUpdateImpossible
		}
		return nil, err
	}

	return entitymap.MapToReception(&reception, nil), nil
}

func (s *receptionStorage) insertNewReception(ctx context.Context, tx *sql.Tx, pvzID int) (*entities.Reception, error) {
	query, args, err := sq.Insert(consts.ReceptionsTable).
		Columns(consts.ReceptionPvzID).
		Values(pvzID).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var reception entities.Reception

	row := tx.QueryRowContext(ctx, query, args...)
	err = row.Scan(&reception.ID, &reception.DateTime, &reception.PvzID, &reception.Status)
	if err != nil {
		return nil, s.helper.CatchPQErrors(err)
	}

	return &reception, nil
}

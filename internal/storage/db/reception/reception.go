package reception

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"pvz-service/internal/entities"
	"pvz-service/internal/enum"
	"pvz-service/internal/models"
	"pvz-service/internal/storage/db/dberrors"
	"pvz-service/internal/storage/db/shared/consts"

	entitymap "pvz-service/internal/mappers/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

type ReceptionDB interface {
	CreateReception(ctx context.Context, createreception *models.CreateReception) (*models.Reception, error)
	CloseReception(ctx context.Context, closeReception *models.CloseReception) (*models.Reception, error)
}

type receptionStorage struct {
	db     *sql.DB
	logger *slog.Logger
}

func New(db *sql.DB, logger *slog.Logger) ReceptionDB {
	return &receptionStorage{
		db:     db,
		logger: logger,
	}
}

func (s *receptionStorage) CreateReception(ctx context.Context, createReception *models.CreateReception) (*models.Reception, error) {
	selectQuery, selArgs, err := sq.Select(consts.ID).
		From(consts.ReceptionsTable).
		Where(sq.Eq{
			consts.ReceptionPvzID:  createReception.PvzID,
			consts.ReceptionStatus: enum.StatusInProgress.String(),
		}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	insertQuery, insArgs, err := sq.Insert(consts.ReceptionsTable).
		Columns(consts.ReceptionPvzID).
		Values(createReception.PvzID).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var reception entities.Reception

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	var plug int

	row := tx.QueryRowContext(ctx, selectQuery, selArgs...)
	selectErr := row.Scan(&plug)
	if selectErr != nil && selectErr == sql.ErrNoRows {
		row = tx.QueryRowContext(ctx, insertQuery, insArgs...)
		err := row.Scan(&reception.ID, &reception.DateTime, &reception.PvzID, &reception.Status)
		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				s.logger.Error("tx rollback error: " + txErr.Error())
			}
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code.Name() == consts.PQInvalidTextRepresentation {
					return nil, errors.Join(dberrors.ErrEnumTypeViolation, err) // For mode specifics in logs.
				}
				if pqErr.Code.Name() == consts.PQForeignKeyViolation {
					return nil, dberrors.ErrForeignKeyViolation
				}
			}
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	if selectErr == nil {
		return nil, dberrors.ErrInsertImpossible
	} else if selectErr != sql.ErrNoRows {
		return nil, selectErr
	}

	return entitymap.MapToReception(&reception, nil), nil
}

func (s *receptionStorage) CloseReception(ctx context.Context, closeReception *models.CloseReception) (*models.Reception, error) {
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

package reception

import (
	"context"
	"database/sql"
	"errors"
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
	db *sql.DB
}

func New(db *sql.DB) ReceptionDB {
	return &receptionStorage{
		db: db,
	}
}

func (s *receptionStorage) CreateReception(ctx context.Context, createReception *models.CreateReception) (*models.Reception, error) {
	selectQuery, selArgs, err := sq.Select(consts.ID).
		From(consts.ReceptionsTable).
		Where(sq.Eq{
			consts.PvzID:  createReception.PvzID,
			consts.Status: enum.StatusInProgress.String(),
		}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	insertQuery, insArgs, err := sq.Insert(consts.ReceptionsTable).
		Columns(consts.PvzID).
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
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code.Name() == consts.PQInvalidTextRepresentation {
					return nil, errors.Join(dberrors.ErrEnumTypeViolation, err) // For mode specifics in logs.
				}
				if pqErr.Code.Name() == consts.PQForeignKeyViolation {
					return nil, dberrors.ErrForeignKeyViolation
				}
			}
		}
	} else {
		err := tx.Commit()
		if err != nil {
			return nil, err
		}

		return nil, dberrors.ErrInsertImpossible
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return entitymap.MapToReception(&reception, nil), nil
}

func (s *receptionStorage) CloseReception(ctx context.Context, closeReception *models.CloseReception) (*models.Reception, error) {
	query, args, err := sq.Update(consts.ReceptionsTable).
		Set(consts.Status, enum.StatusClose.String()).
		Where(sq.Eq{
			consts.PvzID:  closeReception.PvzID,
			consts.Status: enum.StatusInProgress.String(),
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

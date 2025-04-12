package reception

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"pvz-service/internal/entities"
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
	insertQuery, args, err := sq.Insert(consts.ReceptionsTable).
		Columns(consts.PvzID).
		Values(createReception.PvzID).
		Suffix(fmt.Sprintf("RETURNING %s,%s,%s,%s", consts.ID, consts.DateTime, consts.PvzID, consts.Status)).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var reception entities.Reception
	row := s.db.QueryRowContext(ctx, insertQuery, args...)
	err = row.Scan(&reception.ID, &reception.DateTime, &reception.PvzID, &reception.Status)
	if err != nil {
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

	return entitymap.MapToReception(&reception, nil), nil
}

func (s *receptionStorage) CloseReception(ctx context.Context, closeReception *models.CloseReception) (*models.Reception, error) {
	return nil, errors.New("testing plug")
}

package pvz

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"pvz-service/internal/entities"
	entitymap "pvz-service/internal/mappers/entity"
	"pvz-service/internal/models"
	"pvz-service/internal/storage/db/dberrors"
	"pvz-service/internal/storage/db/shared/consts"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

type PvzDB interface {
	CreatePvz(ctx context.Context, pvzCreate *models.CreatePvz) (*models.Pvz, error)
	GetPvzList(ctx context.Context, filters *models.PvzFilterParams) (*models.PvzList, error)
}

type pvzStorage struct {
	db     *sql.DB
	logger *slog.Logger
}

func New(db *sql.DB, logger *slog.Logger) PvzDB {
	return &pvzStorage{
		db:     db,
		logger: logger,
	}
}

func (s *pvzStorage) CreatePvz(ctx context.Context, pvzCreate *models.CreatePvz) (*models.Pvz, error) {
	insertQuery, args, err := sq.Insert(consts.PvzsTable).
		Columns(consts.PvzCity).
		Values(pvzCreate.City).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var pvz entities.Pvz
	row := s.db.QueryRowContext(ctx, insertQuery, args...)
	err = row.Scan(&pvz.ID, &pvz.RegistrationDate, &pvz.City)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code.Name() == consts.PQInvalidTextRepresentation {
				return nil, errors.Join(dberrors.ErrEnumTypeViolation, err)
			}
		}
		return nil, err
	}

	return entitymap.MapToPvz(&pvz), nil
}

func (s *pvzStorage) GetPvzList(ctx context.Context, filters *models.PvzFilterParams) (*models.PvzList, error) {
	return nil, errors.New("testing plug")
}

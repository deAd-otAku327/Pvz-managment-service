package pvz

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"pvz-service/internal/entities"
	entitymap "pvz-service/internal/mappers/entity"
	"pvz-service/internal/models"
	"pvz-service/internal/storage/db/dberrors"
	"pvz-service/internal/storage/db/shared/consts"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

type PvzDB interface {
	CreatePvz(ctx context.Context, pvzCreate *entities.CreatePvz) (*models.Pvz, error)
	GetPvzList(ctx context.Context, filters *entities.PvzFilterParams) (models.PvzList, error)
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

func (s *pvzStorage) CreatePvz(ctx context.Context, pvzCreate *entities.CreatePvz) (*models.Pvz, error) {
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

func (s *pvzStorage) GetPvzList(ctx context.Context, filters *entities.PvzFilterParams) (models.PvzList, error) {
	pvzs, err := s.getPvzsWithPagination(ctx, filters.Page, filters.Limit)
	if err != nil {
		return nil, err
	}

	pvzIDs := make([]int, 0)
	for _, pvz := range pvzs {
		pvzIDs = append(pvzIDs, pvz.ID)
	}

	receptions, err := s.getRelatedReceptionsWithDateInterval(ctx, pvzIDs, filters.StartDate, filters.EndDate)
	if err != nil {
		return nil, err
	}

	receptionIDs := make([]int, 0)
	for _, r := range receptions {
		receptionIDs = append(receptionIDs, r.ID)
	}

	products, err := s.getRelatedProducts(ctx, receptionIDs)
	if err != nil {
		return nil, err
	}

	return entitymap.MapToPvzList(pvzs, receptions, products), nil
}

func (s *pvzStorage) getPvzsWithPagination(ctx context.Context, page, limit int) ([]*entities.Pvz, error) {
	offset := (page - 1) * limit
	query, args, err := sq.Select("*").
		From(consts.PvzsTable).
		OrderBy(consts.ID).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	pvzs := make([]*entities.Pvz, 0)

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pvz := entities.Pvz{}
		err := rows.Scan(&pvz.ID, &pvz.RegistrationDate, &pvz.City)
		if err != nil {
			return nil, err
		}
		pvzs = append(pvzs, &pvz)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return pvzs, nil
}

func (s *pvzStorage) getRelatedReceptionsWithDateInterval(ctx context.Context, pvzsIDs []int, startDate, endDate time.Time) ([]*entities.Reception, error) {
	query, args, err := sq.Select("*").
		From(consts.ReceptionsTable).
		Where(sq.GtOrEq{consts.ReceptionDateTime: startDate}).
		Where(sq.LtOrEq{consts.ReceptionDateTime: endDate}).
		Where(fmt.Sprintf("%s = ANY($3)", consts.ReceptionPvzID)).
		OrderBy(consts.ID).
		PlaceholderFormat(sq.Dollar).ToSql()
	args = append(args, pq.Array(pvzsIDs))

	if err != nil {
		return nil, err
	}
	receptions := make([]*entities.Reception, 0)

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		reception := entities.Reception{}
		err := rows.Scan(&reception.ID, &reception.DateTime, &reception.PvzID, &reception.Status)
		if err != nil {
			return nil, err
		}
		receptions = append(receptions, &reception)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return receptions, nil
}

func (s *pvzStorage) getRelatedProducts(ctx context.Context, receptionIDs []int) ([]*entities.Product, error) {
	query, args, err := sq.Select("*").
		From(consts.ProductsTable).
		Where(fmt.Sprintf("%s = ANY($1)", consts.ProductReceptionID)).
		OrderBy(consts.ID).
		PlaceholderFormat(sq.Dollar).ToSql()
	args = append(args, pq.Array(receptionIDs))

	if err != nil {
		return nil, err
	}
	products := make([]*entities.Product, 0)

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := entities.Product{}
		err := rows.Scan(&product.ID, &product.DateTime, &product.ReceptionID, &product.Type)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return products, nil
}

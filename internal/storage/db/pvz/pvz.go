package pvz

import (
	"context"
	"database/sql"
	"errors"
	"pvz-service/internal/models"
)

type PvzDB interface {
	CreatePvz(ctx context.Context, pvzCreate *models.PvzCreate) (*models.Pvz, error)
	GetPvzList(ctx context.Context, filters *models.PvzFilterParams) (*models.PvzList, error)
}

type pvzStorage struct {
	db *sql.DB
}

func New(db *sql.DB) PvzDB {
	return &pvzStorage{
		db: db,
	}
}

func (s *pvzStorage) CreatePvz(ctx context.Context, pvzCreate *models.PvzCreate) (*models.Pvz, error) {
	return nil, errors.New("testing plug")
}

func (s *pvzStorage) GetPvzList(ctx context.Context, filters *models.PvzFilterParams) (*models.PvzList, error) {
	return nil, errors.New("testing plug")
}

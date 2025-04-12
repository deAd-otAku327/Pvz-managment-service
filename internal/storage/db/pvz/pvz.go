package pvz

import (
	"context"
	"database/sql"
	"errors"
	"pvz-service/internal/entities"
	"pvz-service/internal/models"
)

type PvzDB interface {
	CreatePvz(ctx context.Context, pvzCreate *models.PvzCreate) (*entities.Pvz, error)
	GetPvzList(ctx context.Context, filters *models.PvzFilterParams) ([]entities.Pvz, []entities.Reception, error)
}

type pvzStorage struct {
	db *sql.DB
}

func New(db *sql.DB) PvzDB {
	return &pvzStorage{
		db: db,
	}
}

func (s *pvzStorage) CreatePvz(ctx context.Context, pvzCreate *models.PvzCreate) (*entities.Pvz, error) {
	return nil, errors.New("testing plug")
}

func (s *pvzStorage) GetPvzList(ctx context.Context, filters *models.PvzFilterParams) ([]entities.Pvz, []entities.Reception, error) {
	return nil, nil, errors.New("testing plug")
}

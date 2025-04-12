package reception

import (
	"context"
	"database/sql"
	"errors"
	"pvz-service/internal/entities"
	"pvz-service/internal/models"
)

type ReceptionDB interface {
	CreateReception(ctx context.Context, createreception *models.CreateReception) (*entities.Reception, error)
	CloseReception(ctx context.Context, closeReception *models.CloseReception) (*entities.Reception, error)
}

type receptionStorage struct {
	db *sql.DB
}

func New(db *sql.DB) ReceptionDB {
	return &receptionStorage{
		db: db,
	}
}

func (s *receptionStorage) CloseReception(ctx context.Context, closeReception *models.CloseReception) (*entities.Reception, error) {
	return nil, errors.New("testing plug")
}

func (s *receptionStorage) CreateReception(ctx context.Context, createReception *models.CreateReception) (*entities.Reception, error) {
	return nil, errors.New("testing plug")
}

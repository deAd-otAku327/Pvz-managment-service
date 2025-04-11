package db

import (
	"context"
	"errors"
	"pvz-service/internal/entities"
	"pvz-service/internal/models"
)

func (s *storage) CreatePvz(ctx context.Context, pvzCreate *models.PvzCreate) (*entities.Pvz, error) {
	return nil, errors.New("testing plug")
}

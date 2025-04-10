package db

import (
	"context"
	"errors"
	"pvz-service/internal/models"
)

func (s *storage) CreatePvz(ctx context.Context, city string) (*models.PVZ, error) {
	return nil, errors.New("testing plug")
}

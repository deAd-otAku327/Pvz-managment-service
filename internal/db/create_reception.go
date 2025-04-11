package db

import (
	"context"
	"errors"
	"pvz-service/internal/entities"
	"pvz-service/internal/models"
)

func (s *storage) CreateReception(ctx context.Context, createReception *models.CreateReception) (*entities.Reception, error) {
	return nil, errors.New("testing plug")
}

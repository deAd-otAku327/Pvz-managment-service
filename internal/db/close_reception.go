package db

import (
	"context"
	"errors"
	"pvz-service/internal/entities"
	"pvz-service/internal/models"
)

func (s *storage) CloseReception(ctx context.Context, closeReception *models.CloseReception) (*entities.Reception, error) {
	return nil, errors.New("testing plug")
}

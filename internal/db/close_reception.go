package db

import (
	"context"
	"errors"
	"pvz-service/internal/models"
)

func (s *storage) CloseReception(ctx context.Context, pvzID int) (*models.Reception, error) {
	return nil, errors.New("testing plug")
}

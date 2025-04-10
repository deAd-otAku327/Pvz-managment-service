package db

import (
	"context"
	"errors"
	"pvz-service/internal/models"
)

func (s *storage) AddProduct(ctx context.Context, pType string, pvzID int) (*models.Product, error) {
	return nil, errors.New("testing plug")
}

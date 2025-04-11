package db

import (
	"context"
	"errors"
	"pvz-service/internal/entities"
	"pvz-service/internal/models"
)

func (s *storage) AddProduct(ctx context.Context, addProduct *models.AddProduct) (*entities.Product, error) {
	return nil, errors.New("testing plug")
}

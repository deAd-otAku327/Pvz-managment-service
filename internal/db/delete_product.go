package db

import (
	"context"
	"errors"
	"pvz-service/internal/models"
)

func (s *storage) DeleteProduct(ctx context.Context, deleteProduct *models.DeleteProduct) error {
	return errors.New("testing plug")
}

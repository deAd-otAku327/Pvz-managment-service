package product

import (
	"context"
	"database/sql"
	"errors"
	"pvz-service/internal/models"
)

type ProductDB interface {
	AddProduct(ctx context.Context, addProduct *models.AddProduct) (*models.Product, error)
	DeleteProduct(ctx context.Context, deleteProduct *models.DeleteProduct) error
}

type productStorage struct {
	db *sql.DB
}

func New(db *sql.DB) ProductDB {
	return &productStorage{
		db: db,
	}
}

func (s *productStorage) AddProduct(ctx context.Context, addProduct *models.AddProduct) (*models.Product, error) {
	return nil, errors.New("testing plug")
}

func (s *productStorage) DeleteProduct(ctx context.Context, deleteProduct *models.DeleteProduct) error {
	return errors.New("testing plug")
}

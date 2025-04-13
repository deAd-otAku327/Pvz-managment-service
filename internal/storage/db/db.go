package db

import (
	"context"
	"database/sql"
	"log/slog"
	"pvz-service/internal/config"
	"pvz-service/internal/models"
	"pvz-service/internal/storage/db/product"
	"pvz-service/internal/storage/db/pvz"
	"pvz-service/internal/storage/db/reception"

	_ "github.com/lib/pq"
)

type DB interface {
	pvz.PvzDB
	reception.ReceptionDB
	product.ProductDB
}

type storage struct {
	pvzStorage       pvz.PvzDB
	receptionStorage reception.ReceptionDB
	productStorage   product.ProductDB
}

func New(cfg config.DBConn, logger *slog.Logger) (DB, error) {
	database, err := sql.Open("postgres", cfg.URL)
	if err != nil {
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}

	database.SetMaxOpenConns(cfg.MaxOpenConns)

	return &storage{
		pvzStorage:       pvz.New(database, logger),
		receptionStorage: reception.New(database, logger),
		productStorage:   product.New(database, logger),
	}, nil
}

func (s *storage) CreatePvz(ctx context.Context, pvzCreate *models.CreatePvz) (*models.Pvz, error) {
	return s.pvzStorage.CreatePvz(ctx, pvzCreate)
}

func (s *storage) GetPvzList(ctx context.Context, filters *models.PvzFilterParams) (models.PvzList, error) {
	return s.pvzStorage.GetPvzList(ctx, filters)
}

func (s *storage) CreateReception(ctx context.Context, createReception *models.CreateReception) (*models.Reception, error) {
	return s.receptionStorage.CreateReception(ctx, createReception)
}

func (s *storage) CloseReception(ctx context.Context, closeReception *models.CloseReception) (*models.Reception, error) {
	return s.receptionStorage.CloseReception(ctx, closeReception)
}

func (s *storage) AddProduct(ctx context.Context, addProduct *models.AddProduct) (*models.Product, error) {
	return s.productStorage.AddProduct(ctx, addProduct)
}

func (s *storage) DeleteLastProduct(ctx context.Context, deleteProduct *models.DeleteProduct) error {
	return s.productStorage.DeleteLastProduct(ctx, deleteProduct)
}

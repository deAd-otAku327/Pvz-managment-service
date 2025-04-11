package db

import (
	"context"
	"database/sql"
	"pvz-service/internal/config"
	"pvz-service/internal/entities"
	"pvz-service/internal/models"

	_ "github.com/lib/pq"
)

type DB interface {
	CreatePvz(ctx context.Context, pvzCreate *models.PvzCreate) (*entities.Pvz, error)
	GetPvzList(ctx context.Context, filters *models.PvzFilterParams) ([]entities.Pvz, []entities.Reception, error)
	CreateReception(ctx context.Context, createreception *models.CreateReception) (*entities.Reception, error)
	CloseReception(ctx context.Context, closeReception *models.CloseReception) (*entities.Reception, error)
	AddProduct(ctx context.Context, addProduct *models.AddProduct) (*entities.Product, error)
	DeleteProduct(ctx context.Context, deleteProduct *models.DeleteProduct) error
}

type storage struct {
	db *sql.DB
}

func New(cfg config.DBConn) (DB, error) {
	database, err := sql.Open("postgres", cfg.URL)
	if err != nil {
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}

	database.SetMaxOpenConns(cfg.MaxOpenConns)

	return &storage{db: database}, nil
}

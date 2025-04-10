package db

import (
	"context"
	"database/sql"
	"pvz-service/internal/config"
	"pvz-service/internal/models"

	_ "github.com/lib/pq"
)

type DB interface {
	CreatePvz(ctx context.Context, city string) (*models.PVZ, error)
	GetPvzList(ctx context.Context, filters *models.FilterParams) (*models.SummaryInfo, error)
	CreateReception(ctx context.Context, pvzID int) (*models.Reception, error)
	AddProduct(ctx context.Context, pType string, pvzID int) (*models.Product, error)
	CloseReception(ctx context.Context, pvzID int) (*models.Reception, error)
	DeleteProduct(ctx context.Context, pvzID int) error
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

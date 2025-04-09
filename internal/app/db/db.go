package db

import (
	"database/sql"
	"pvz-service/internal/app/config"

	_ "github.com/lib/pq"
)

type DB interface {
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

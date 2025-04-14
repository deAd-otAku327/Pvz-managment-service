package integration

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"pvz-service/internal/app"
	"pvz-service/internal/config"
	"pvz-service/internal/enum"
	"pvz-service/pkg/logger"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/ory/dockertest/v3"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var server *app.Server

var roles = map[string]string{
	enum.Employye.String():  "",
	enum.Moderator.String(): "",
}

func initTestServer(dbPort string) error {
	cfg := &config.Config{
		Server: config.Server{
			JWTKey:       "testing",
			Host:         "localhost",
			Port:         "8081",
			ResponseTime: 100 * time.Millisecond,
			RPS:          1000,
			LogLevel:     "error",
		},
		DBConn: config.DBConn{
			URL: fmt.Sprintf("postgres://postgres:qwerty123@localhost:%s/pvz-db?sslmode=disable", dbPort),
		},
	}

	logger, err := logger.New(os.Stdout, cfg.LogLevel)
	if err != nil {
		return err
	}

	server, err = app.New(cfg, logger)
	if err != nil {
		return err
	}

	return nil
}

func createPostgresDB(pool *dockertest.Pool, name string) (*dockertest.Resource, error) {
	log.Printf("Postgres starting...")
	res, err := pool.RunWithOptions(&dockertest.RunOptions{
		Name: name,
		Env: []string{
			"PGUSER=postgres",
			"POSTGRES_PASSWORD=qwerty123",
			"POSTGRES_DB=pvz-db",
		},
		Repository: "postgres",
		Tag:        "latest",
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Postgres started")

	dbPath := fmt.Sprintf("postgres://postgres:qwerty123@localhost:%s/pvz-db?sslmode=disable", res.GetPort("5432/tcp"))

	var db *sql.DB
	if err := pool.Retry(func() error {
		log.Println("Checking postgres connection...")
		db, err = sql.Open("postgres", dbPath)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		return nil, err
	}

	log.Println("Postgres connection established")

	migrationsPath := "file://../../internal/storage/migrations"

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		return nil, err
	}

	if err := m.Up(); err != nil {
		return nil, err
	}

	return res, nil
}

// func createServer(pool *dockertest.Pool, network *dockertest.Network, name string) (*dockertest.Resource, error) {
// 	log.Printf("Server starting...")
// 	res, err := pool.BuildAndRunWithBuildOptions(&dockertest.BuildOptions{
// 		Dockerfile: "Dockerfile",
// 		ContextDir: "../../",
// 	}, &dockertest.RunOptions{
// 		Name:         name,
// 		Env:          []string{"ENV=integration", "JWTKEY=testing"},
// 		ExposedPorts: []string{"8080"},
// 		NetworkID:    network.Network.ID,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	serverAddress = res.GetHostPort("8080/tcp")
// 	log.Printf("Server started on %v", serverAddress)

// 	return res, nil
// }

func teardown(pool *dockertest.Pool, resource *dockertest.Resource) error {
	if err := pool.Purge(resource); err != nil {
		return err
	}

	return nil
}

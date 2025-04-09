package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"pvz-service/internal/app/config"
	"pvz-service/internal/app/controller"
	"pvz-service/internal/app/db"
	"pvz-service/internal/app/service"
	"pvz-service/internal/tokenizer"
	"pvz-service/pkg/cryptor"
	"pvz-service/pkg/middleware"

	"github.com/gorilla/mux"
)

const AppName = "pvz-management-service"

type Server struct {
	cfg    *config.Config
	server *http.Server
}

func New(cfg *config.Config, logger *slog.Logger) (*Server, error) {
	cryptor := cryptor.New()
	tokenizer := tokenizer.New(AppName, cfg.JWTKey)

	storage, err := db.New(cfg.DBConn)
	if err != nil {
		return nil, err
	}

	service := service.New(storage, logger, cryptor, tokenizer)

	controller := controller.New(service)

	router := mux.NewRouter()

	router.Use(middleware.RpsLimit(cfg.RPS))
	router.Use(middleware.Logging(logger))
	router.Use(middleware.ResponseTimeLimit(cfg.ResponseTime))

	router.HandleFunc("/dummyLogin", controller.DummyLogin()).Methods(http.MethodPost)

	return &Server{
		cfg: cfg,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
			Handler: router,
		},
	}, nil
}

func (s *Server) Run() error {
	slog.Info("server starting on", slog.String("address", s.server.Addr))
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown() error {
	return s.server.Shutdown(context.Background())
}

package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"pvz-service/internal/app/config"

	"github.com/gorilla/mux"
)

type Server struct {
	cfg    *config.Config
	server *http.Server
}

func New(cfg *config.Config, logger *slog.Logger) (*Server, error) {

	router := mux.NewRouter()

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

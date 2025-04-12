package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"pvz-service/internal/config"
	"pvz-service/internal/controller"
	"pvz-service/internal/db"
	"pvz-service/internal/enum"
	"pvz-service/internal/middleware"
	"pvz-service/internal/service"
	"pvz-service/internal/tokenizer"
	"pvz-service/pkg/cryptor"

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
	router.HandleFunc("/register", controller.Register()).Methods(http.MethodPost)
	router.HandleFunc("/login", controller.Login()).Methods(http.MethodPost)

	modRouter := router.PathPrefix("").Subrouter()
	modRouter.Use(middleware.AuthOnRoles(tokenizer, map[string]struct{}{
		enum.Moderator.String(): {},
	}))

	modRouter.HandleFunc("/pvz", controller.CreatePvz()).Methods(http.MethodPost)

	empRouter := router.PathPrefix("").Subrouter()
	empRouter.Use(middleware.AuthOnRoles(tokenizer, map[string]struct{}{
		enum.Employye.String(): {},
	}))

	empRouter.HandleFunc("/receptions", controller.CreateReception()).Methods(http.MethodPost)
	empRouter.HandleFunc("/products", controller.AddProduct()).Methods(http.MethodPost)
	empRouter.HandleFunc("/pvz/{pvzId:[1-9][0-9]*}/close_last_reception", controller.CloseReception()).Methods(http.MethodPost)
	empRouter.HandleFunc("/pvz/{pvzId:[1-9][0-9]*}/delete_last_product", controller.DeleteProduct()).Methods(http.MethodPost)

	modAndEmpRouter := router.PathPrefix("").Subrouter()
	modAndEmpRouter.Use(middleware.AuthOnRoles(tokenizer, map[string]struct{}{
		enum.Employye.String():  {},
		enum.Moderator.String(): {},
	}))

	modAndEmpRouter.HandleFunc("/pvz", controller.GetPvzList()).Methods(http.MethodGet)

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

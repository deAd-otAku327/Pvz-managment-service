package service

import (
	"context"
	"log/slog"
	"pvz-service/internal/db"
	"pvz-service/internal/models"
	"pvz-service/internal/service/auth"
	"pvz-service/internal/service/product"
	"pvz-service/internal/service/pvz"
	"pvz-service/internal/service/reception"
	"pvz-service/internal/tokenizer"
	"pvz-service/pkg/cryptor"
	"pvz-service/pkg/werrors"
)

type Service interface {
	auth.AuthService
	pvz.PvzService
	reception.ReceptionService
	product.ProductService
}

type service struct {
	authService      auth.AuthService
	productService   product.ProductService
	receptionService reception.ReceptionService
	pvzService       pvz.PvzService
}

func New(storage db.DB, log *slog.Logger, cryptor cryptor.Cryptor, tok tokenizer.Tokenizer) Service {
	return &service{
		authService:      auth.New(storage, log, cryptor, tok),
		productService:   product.New(storage, log),
		receptionService: reception.New(storage, log),
		pvzService:       pvz.New(storage, log),
	}
}

func (s *service) DummyLogin(ctx context.Context, role string) (*string, werrors.Werror) {
	return s.authService.DummyLogin(ctx, role)
}

func (s *service) AddProduct(ctx context.Context, pType string, pvzID int) (*models.Product, werrors.Werror) {
	return s.productService.AddProduct(ctx, pType, pvzID)
}

func (s *service) DeleteProduct(ctx context.Context, id string) werrors.Werror {
	return s.productService.DeleteProduct(ctx, id)
}

func (s *service) CreateReception(ctx context.Context, pvzID int) (*models.Reception, werrors.Werror) {
	return s.receptionService.CreateReception(ctx, pvzID)
}

func (s *service) CloseReception(ctx context.Context, pvzID string) (*models.Reception, werrors.Werror) {
	return s.receptionService.CloseReception(ctx, pvzID)
}

func (s *service) CreatePvz(ctx context.Context, city string) (*models.PVZ, werrors.Werror) {
	return s.pvzService.CreatePvz(ctx, city)
}

func (s *service) GetPvzList(ctx context.Context, startDate, endDate, page, limit string) (*models.SummaryInfo, werrors.Werror) {
	return s.pvzService.GetPvzList(ctx, startDate, endDate, page, limit)
}

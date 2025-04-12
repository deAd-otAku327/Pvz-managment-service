package service

import (
	"context"
	"log/slog"
	"pvz-service/internal/dto"
	"pvz-service/internal/service/auth"
	"pvz-service/internal/service/product"
	"pvz-service/internal/service/pvz"
	"pvz-service/internal/service/reception"
	"pvz-service/internal/storage/db"
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

func (s *service) AddProduct(ctx context.Context, request *dto.AddProductRequestDTO) (*dto.ProductResponseDTO, werrors.Werror) {
	return s.productService.AddProduct(ctx, request)
}

func (s *service) DeleteProduct(ctx context.Context, request *dto.DeleteProductRequestDTO) werrors.Werror {
	return s.productService.DeleteProduct(ctx, request)
}

func (s *service) CreateReception(ctx context.Context, request *dto.CreateReceptionRequestDTO) (*dto.ReceptionResponseDTO, werrors.Werror) {
	return s.receptionService.CreateReception(ctx, request)
}

func (s *service) CloseReception(ctx context.Context, request *dto.CloseReceptionRequestDTO) (*dto.ReceptionResponseDTO, werrors.Werror) {
	return s.receptionService.CloseReception(ctx, request)
}

func (s *service) CreatePvz(ctx context.Context, request *dto.CreatePvzRequestDTO) (*dto.PvzResponseDTO, werrors.Werror) {
	return s.pvzService.CreatePvz(ctx, request)
}

func (s *service) GetPvzList(ctx context.Context, request *dto.PvzFilterParamsDTO) (*dto.GetPvzListResponseDTO, werrors.Werror) {
	return s.pvzService.GetPvzList(ctx, request)
}

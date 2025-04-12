package product

import (
	"context"
	"log/slog"
	"net/http"
	"pvz-service/internal/apperrors"
	"pvz-service/internal/dto"
	dtomap "pvz-service/internal/mappers/dto"
	modelmap "pvz-service/internal/mappers/model"
	"pvz-service/internal/storage/db"
	"pvz-service/pkg/werrors"
)

type ProductService interface {
	AddProduct(ctx context.Context, request *dto.AddProductRequestDTO) (*dto.ProductResponseDTO, werrors.Werror)
	DeleteProduct(ctx context.Context, request *dto.DeleteProductRequestDTO) werrors.Werror
}

type productService struct {
	storage db.DB
	logger  *slog.Logger
}

func New(storage db.DB, logger *slog.Logger) ProductService {
	return &productService{
		storage: storage,
		logger:  logger,
	}
}

func (s *productService) DeleteProduct(ctx context.Context, request *dto.DeleteProductRequestDTO) werrors.Werror {
	deleteProduct := dtomap.MapToDeleteProduct(request)
	err := deleteProduct.Validate()
	if err != nil {
		return werrors.New(err, http.StatusBadRequest)
	}

	err = s.storage.DeleteProduct(ctx, deleteProduct)
	if err != nil {
		s.logger.Error("delete product: " + err.Error())
		return werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	return nil
}

func (s *productService) AddProduct(ctx context.Context, request *dto.AddProductRequestDTO) (*dto.ProductResponseDTO, werrors.Werror) {
	addProduct := dtomap.MapToAddProduct(request)
	err := addProduct.Validate()
	if err != nil {
		return nil, werrors.New(err, http.StatusBadRequest)
	}

	product, err := s.storage.AddProduct(ctx, addProduct)
	if err != nil {
		s.logger.Error("add product: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	err = product.Validate()
	if err != nil {
		s.logger.Error("add product response data invalid, DB inconsistency detected: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToProductResponse(product), nil
}

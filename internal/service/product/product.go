package product

import (
	"context"
	"log/slog"
	"net/http"
	"pvz-service/internal/apperrors"
	"pvz-service/internal/dto"

	modelmap "pvz-service/internal/mappers/model"
	"pvz-service/internal/models"
	"pvz-service/internal/storage/db"
	"pvz-service/internal/storage/db/dberrors"
	"pvz-service/pkg/werrors"
)

type ProductService interface {
	AddProduct(ctx context.Context, request *models.AddProduct) (*dto.ProductResponseDTO, werrors.Werror)
	DeleteProduct(ctx context.Context, request *models.DeleteProduct) werrors.Werror
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

func (s *productService) DeleteProduct(ctx context.Context, deleteProduct *models.DeleteProduct) werrors.Werror {
	err := deleteProduct.Validate()
	if err != nil {
		return werrors.New(err, http.StatusBadRequest)
	}

	err = s.storage.DeleteLastProduct(ctx, modelmap.MapToDeleteProduct(deleteProduct))
	if err != nil {
		if err == dberrors.ErrNothingToDelete {
			return werrors.New(apperrors.ErrNoProductsInCurrReception, http.StatusBadRequest)
		}
		if err == dberrors.ErrDeleteImpossible {
			return werrors.New(apperrors.ErrReceptionIsNotOpened, http.StatusBadRequest)
		}
		s.logger.Error("delete product: " + err.Error())
		return werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	return nil
}

func (s *productService) AddProduct(ctx context.Context, addProduct *models.AddProduct) (*dto.ProductResponseDTO, werrors.Werror) {
	err := addProduct.Validate()
	if err != nil {
		return nil, werrors.New(err, http.StatusBadRequest)
	}

	product, err := s.storage.AddProduct(ctx, modelmap.MapToAddProduct(addProduct))
	if err != nil {
		if err == dberrors.ErrInsertImpossible {
			return nil, werrors.New(apperrors.ErrReceptionIsNotOpened, http.StatusBadRequest)
		}
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

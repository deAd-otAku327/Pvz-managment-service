package product

import (
	"context"
	"log/slog"
	"net/http"
	"pvz-service/internal/apperrors"
	"pvz-service/internal/db"
	"pvz-service/internal/enum"
	"pvz-service/internal/models"
	"pvz-service/pkg/werrors"
	"strconv"
)

type ProductService interface {
	AddProduct(ctx context.Context, pType string, pvzID int) (*models.Product, werrors.Werror)
	DeleteProduct(ctx context.Context, id string) werrors.Werror
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

func (s *productService) DeleteProduct(ctx context.Context, id string) werrors.Werror {
	pvzId, err := strconv.Atoi(id)
	if err != nil {
		return werrors.New(apperrors.ErrInvalidPvzID, http.StatusBadRequest)
	}

	err = s.storage.DeleteProduct(ctx, pvzId)
	if err != nil {
		s.logger.Error("delete product: " + err.Error())
		return werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	return nil
}

func (s *productService) AddProduct(ctx context.Context, pType string, pvzID int) (*models.Product, werrors.Werror) {
	if !enum.CheckProductType(pType) {
		return nil, werrors.New(apperrors.ErrInvalidProductType, http.StatusBadRequest)
	}

	product, err := s.storage.AddProduct(ctx, pType, pvzID)
	if err != nil {
		s.logger.Error("add product: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	return product, nil
}

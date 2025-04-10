package service

import (
	"context"
	"net/http"
	"pvz-service/internal/models"
	"pvz-service/pkg/werrors"
)

func (s *pvzService) AddProduct(ctx context.Context, pType string, pvzID int) (*models.Product, werrors.Werror) {
	if !models.CheckProductType(pType) {
		return nil, werrors.New(errInvalidProductType, http.StatusBadRequest)
	}

	product, err := s.storage.AddProduct(ctx, pType, pvzID)
	if err != nil {
		s.logger.Error("add product: " + err.Error())
		return nil, werrors.New(errSmthWentWrong, http.StatusInternalServerError)
	}

	return product, nil
}

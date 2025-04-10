package service

import (
	"context"
	"net/http"
	"pvz-service/internal/models"
	"pvz-service/pkg/werrors"
)

func (s *pvzService) CreatePvz(ctx context.Context, city string) (*models.PVZ, werrors.Werror) {
	if !models.CheckCity(city) {
		return nil, werrors.New(errInvalidCity, http.StatusBadRequest)
	}

	pvz, err := s.storage.CreatePvz(ctx, city)
	if err != nil {
		s.logger.Error("create pvz: " + err.Error())
		return nil, werrors.New(errSmthWentWrong, http.StatusInternalServerError)
	}

	return pvz, nil
}

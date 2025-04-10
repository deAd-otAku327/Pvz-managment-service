package service

import (
	"context"
	"net/http"
	"pvz-service/internal/models"
	"pvz-service/pkg/werrors"
)

func (s *pvzService) CreateReception(ctx context.Context, pvzID int) (*models.Reception, werrors.Werror) {
	reception, err := s.storage.CreateReception(ctx, pvzID)
	if err != nil {
		s.logger.Error("create reception: " + err.Error())
		return nil, werrors.New(errSmthWentWrong, http.StatusInternalServerError)
	}

	return reception, nil
}

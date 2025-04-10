package service

import (
	"context"
	"net/http"
	"pvz-service/internal/models"
	"pvz-service/pkg/werrors"
	"strconv"
)

func (s *pvzService) CloseReception(ctx context.Context, id string) (*models.Reception, werrors.Werror) {
	pvzId, err := strconv.Atoi(id) // Some extra validation after reqexps in routing.
	if err != nil {
		return nil, werrors.New(errInvalidPvzID, http.StatusBadRequest)
	}

	reception, err := s.storage.CloseReception(ctx, pvzId)
	if err != nil {
		s.logger.Error("close reception: " + err.Error())
		return nil, werrors.New(errSmthWentWrong, http.StatusInternalServerError)
	}

	return reception, nil
}

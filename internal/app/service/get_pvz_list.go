package service

import (
	"context"
	"net/http"
	"pvz-service/internal/models"
	"pvz-service/pkg/werrors"
)

func (s *pvzService) GetPvzList(ctx context.Context, startDate, endDate, page, limit string) (*models.SummaryInfo, werrors.Werror) {

	params, err := models.NewFilterParams(
		models.WithStartDate(startDate),
		models.WithEndDate(endDate),
		models.WithPage(page),
		models.WithLimit(limit),
	)
	if err != nil {
		return nil, werrors.New(err, http.StatusBadRequest)
	}

	summaryInfo, err := s.storage.GetPvzList(ctx, params)
	if err != nil {
		s.logger.Error("get pvz listing: " + err.Error())
		return nil, werrors.New(errSmthWentWrong, http.StatusInternalServerError)
	}

	return summaryInfo, nil
}

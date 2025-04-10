package pvz

import (
	"context"
	"log/slog"
	"net/http"
	"pvz-service/internal/apperrors"
	"pvz-service/internal/db"
	"pvz-service/internal/enum"
	"pvz-service/internal/models"
	"pvz-service/pkg/werrors"
)

type PvzService interface {
	CreatePvz(ctx context.Context, city string) (*models.PVZ, werrors.Werror)
	GetPvzList(ctx context.Context, startDate, endDate, page, limit string) (*models.SummaryInfo, werrors.Werror)
}

type pvzService struct {
	storage db.DB
	logger  *slog.Logger
}

func New(storage db.DB, logger *slog.Logger) PvzService {
	return &pvzService{
		storage: storage,
		logger:  logger,
	}
}

func (s *pvzService) GetPvzList(ctx context.Context, startDate, endDate, page, limit string) (*models.SummaryInfo, werrors.Werror) {
	params, err := NewFilterParams(
		WithStartDate(startDate),
		WithEndDate(endDate),
		WithPage(page),
		WithLimit(limit),
	)
	if err != nil {
		return nil, werrors.New(err, http.StatusBadRequest)
	}

	summaryInfo, err := s.storage.GetPvzList(ctx, params)
	if err != nil {
		s.logger.Error("get pvz listing: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	return summaryInfo, nil
}

func (s *pvzService) CreatePvz(ctx context.Context, city string) (*models.PVZ, werrors.Werror) {
	if !enum.CheckCity(city) {
		return nil, werrors.New(apperrors.ErrInvalidCity, http.StatusBadRequest)
	}

	pvz, err := s.storage.CreatePvz(ctx, city)
	if err != nil {
		s.logger.Error("create pvz: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	return pvz, nil
}

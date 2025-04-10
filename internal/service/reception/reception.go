package reception

import (
	"context"
	"log/slog"
	"net/http"
	"pvz-service/internal/apperrors"
	"pvz-service/internal/db"
	"pvz-service/internal/models"
	"pvz-service/pkg/werrors"
	"strconv"
)

type ReceptionService interface {
	CreateReception(ctx context.Context, pvzID int) (*models.Reception, werrors.Werror)
	CloseReception(ctx context.Context, id string) (*models.Reception, werrors.Werror)
}

type receptionService struct {
	storage db.DB
	logger  *slog.Logger
}

func New(storage db.DB, logger *slog.Logger) ReceptionService {
	return &receptionService{
		storage: storage,
		logger:  logger,
	}
}

func (s *receptionService) CreateReception(ctx context.Context, pvzID int) (*models.Reception, werrors.Werror) {
	reception, err := s.storage.CreateReception(ctx, pvzID)
	if err != nil {
		s.logger.Error("create reception: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	return reception, nil
}

func (s *receptionService) CloseReception(ctx context.Context, id string) (*models.Reception, werrors.Werror) {
	pvzId, err := strconv.Atoi(id) // Some extra validation after reqexps in routing.
	if err != nil {
		return nil, werrors.New(apperrors.ErrInvalidPvzID, http.StatusBadRequest)
	}

	reception, err := s.storage.CloseReception(ctx, pvzId)
	if err != nil {
		s.logger.Error("close reception: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	return reception, nil
}

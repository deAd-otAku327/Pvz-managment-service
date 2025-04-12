package pvz

import (
	"context"
	"log/slog"
	"net/http"
	"pvz-service/internal/apperrors"
	"pvz-service/internal/dto"
	modelmap "pvz-service/internal/mappers/model"
	"pvz-service/internal/models"
	"pvz-service/internal/storage/db"
	"pvz-service/pkg/werrors"
)

type PvzService interface {
	CreatePvz(ctx context.Context, createPvz *models.CreatePvz) (*dto.PvzResponseDTO, werrors.Werror)
	GetPvzList(ctx context.Context, pvzFilterParams *models.PvzFilterParams) (*dto.GetPvzListResponseDTO, werrors.Werror)
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

func (s *pvzService) GetPvzList(ctx context.Context, pvzFilterParams *models.PvzFilterParams) (*dto.GetPvzListResponseDTO, werrors.Werror) {
	err := pvzFilterParams.Validate()
	if err != nil {
		return nil, werrors.New(err, http.StatusBadRequest)
	}

	pvzList, err := s.storage.GetPvzList(ctx, pvzFilterParams)
	if err != nil {
		s.logger.Error("get pvz listing: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	err = pvzList.Validate()
	if err != nil {
		s.logger.Error("get pvz listing response data invalid, DB inconsistency detected: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToGetPvzListResponse(pvzList), nil
}

func (s *pvzService) CreatePvz(ctx context.Context, createPvz *models.CreatePvz) (*dto.PvzResponseDTO, werrors.Werror) {
	err := createPvz.Validate()
	if err != nil {
		return nil, werrors.New(err, http.StatusBadRequest)
	}

	pvz, err := s.storage.CreatePvz(ctx, createPvz)
	if err != nil {
		s.logger.Error("create pvz: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	err = pvz.Validate()
	if err != nil {
		s.logger.Error("create pvz response data invalid, DB inconsistency detected: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToPvzResponse(pvz), nil
}

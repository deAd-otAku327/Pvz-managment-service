package pvz

import (
	"context"
	"log/slog"
	"net/http"
	"pvz-service/internal/apperrors"
	"pvz-service/internal/db"
	"pvz-service/internal/dto"
	dtomap "pvz-service/internal/mappers/dto"
	entitymap "pvz-service/internal/mappers/entity"
	"pvz-service/pkg/werrors"
)

type PvzService interface {
	CreatePvz(ctx context.Context, request *dto.CreatePvzRequestDTO) (*dto.PvzResponseDTO, werrors.Werror)
	GetPvzList(ctx context.Context, request *dto.PvzFilterParamsDTO) (*dto.GetPvzListResponseDTO, werrors.Werror)
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

func (s *pvzService) GetPvzList(ctx context.Context, request *dto.PvzFilterParamsDTO) (*dto.GetPvzListResponseDTO, werrors.Werror) {
	params := dtomap.MapToPvzFilterParams(request)
	err := params.Validate()
	if err != nil {
		return nil, werrors.New(err, http.StatusBadRequest)
	}

	pvzs, receptions, err := s.storage.GetPvzList(ctx, params)
	if err != nil {
		s.logger.Error("get pvz listing: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	return entitymap.MapToGetPvzListResponse(pvzs, receptions), nil
}

func (s *pvzService) CreatePvz(ctx context.Context, request *dto.CreatePvzRequestDTO) (*dto.PvzResponseDTO, werrors.Werror) {
	pvzCreate := dtomap.MapToPvzCreate(request)
	err := pvzCreate.Validate()
	if err != nil {
		return nil, werrors.New(err, http.StatusBadRequest)
	}

	pvz, err := s.storage.CreatePvz(ctx, pvzCreate)
	if err != nil {
		s.logger.Error("create pvz: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	return entitymap.MapToPvzResponse(pvz), nil
}

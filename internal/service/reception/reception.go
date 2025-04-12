package reception

import (
	"context"
	"log/slog"
	"net/http"
	"pvz-service/internal/apperrors"
	"pvz-service/internal/dto"
	dtomap "pvz-service/internal/mappers/dto"
	modelmap "pvz-service/internal/mappers/model"
	"pvz-service/internal/storage/db"
	"pvz-service/pkg/werrors"
)

type ReceptionService interface {
	CreateReception(ctx context.Context, request *dto.CreateReceptionRequestDTO) (*dto.ReceptionResponseDTO, werrors.Werror)
	CloseReception(ctx context.Context, request *dto.CloseReceptionRequestDTO) (*dto.ReceptionResponseDTO, werrors.Werror)
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

func (s *receptionService) CreateReception(ctx context.Context, request *dto.CreateReceptionRequestDTO) (*dto.ReceptionResponseDTO, werrors.Werror) {
	createReception := dtomap.MapToCreateReception(request)
	err := createReception.Validate()
	if err != nil {
		return nil, werrors.New(err, http.StatusBadRequest)
	}

	reception, err := s.storage.CreateReception(ctx, createReception)
	if err != nil {
		s.logger.Error("create reception: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	err = reception.Validate()
	if err != nil {
		s.logger.Error("create reception response data invalid, DB inconsistency detected: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToReceptionResponse(reception), nil
}

func (s *receptionService) CloseReception(ctx context.Context, request *dto.CloseReceptionRequestDTO) (*dto.ReceptionResponseDTO, werrors.Werror) {
	closeReception := dtomap.MapToCloseReception(request)
	err := closeReception.Validate()
	if err != nil {
		return nil, werrors.New(err, http.StatusBadRequest)
	}

	reception, err := s.storage.CloseReception(ctx, closeReception)
	if err != nil {
		s.logger.Error("close reception: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	err = reception.Validate()
	if err != nil {
		s.logger.Error("close reception response data invalid, DB inconsistency detected: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToReceptionResponse(reception), nil
}

package reception

import (
	"context"
	"log/slog"
	"net/http"
	"pvz-service/internal/apperrors"
	"pvz-service/internal/dto"
	"pvz-service/internal/entities"
	dtomap "pvz-service/internal/mappers/dto"
	"pvz-service/internal/storage/db"
	"pvz-service/pkg/werrors"
)

type ReceptionService interface {
	CreateReception(ctx context.Context, request *dto.CreateReceptionRequestDTO) (*entities.Reception, werrors.Werror)
	CloseReception(ctx context.Context, request *dto.CloseReceptionRequestDTO) (*entities.Reception, werrors.Werror)
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

func (s *receptionService) CreateReception(ctx context.Context, request *dto.CreateReceptionRequestDTO) (*entities.Reception, werrors.Werror) {
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

	return reception, nil
}

func (s *receptionService) CloseReception(ctx context.Context, request *dto.CloseReceptionRequestDTO) (*entities.Reception, werrors.Werror) {
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

	return reception, nil
}

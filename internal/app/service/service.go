package service

import (
	"context"
	"errors"
	"log/slog"
	"pvz-service/internal/app/db"
	"pvz-service/internal/models"
	"pvz-service/internal/tokenizer"
	"pvz-service/pkg/cryptor"
	"pvz-service/pkg/werrors"
)

var (
	errSmthWentWrong      = errors.New("something went wrong")
	errInvalidPvzID       = errors.New("invalid pvz id provided")
	errInvalidRole        = errors.New("invalid role provided")
	errInvalidCity        = errors.New("invalid city provided")
	errInvalidProductType = errors.New("invalid product type provided")
)

type PvzService interface {
	DummyLogin(ctx context.Context, role string) (*string, werrors.Werror)
	CreatePvz(ctx context.Context, city string) (*models.PVZ, werrors.Werror)
	GetPvzList(ctx context.Context, startDate, endDate, page, limit string) (*models.SummaryInfo, werrors.Werror)
	CreateReception(ctx context.Context, pvzID int) (*models.Reception, werrors.Werror)
	AddProduct(ctx context.Context, pType string, pvzID int) (*models.Product, werrors.Werror)
	CloseReception(ctx context.Context, pvzID string) (*models.Reception, werrors.Werror)
	DeleteProduct(ctx context.Context, pvzID string) werrors.Werror
}

type pvzService struct {
	storage   db.DB
	logger    *slog.Logger
	cryptor   cryptor.Cryptor
	tokenizer tokenizer.Tokenizer
}

func New(storage db.DB, log *slog.Logger, cryptor cryptor.Cryptor, tok tokenizer.Tokenizer) PvzService {
	return &pvzService{
		storage:   storage,
		logger:    log,
		cryptor:   cryptor,
		tokenizer: tok,
	}
}

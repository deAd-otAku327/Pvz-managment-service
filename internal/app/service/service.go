package service

import (
	"context"
	"errors"
	"log/slog"
	"pvz-service/internal/app/db"
	"pvz-service/internal/tokenizer"
	"pvz-service/pkg/cryptor"
	"pvz-service/pkg/werrors"
)

var (
	errSmthWentWrong = errors.New("something went wrong")
	errInvalidRole   = errors.New("invalid role provided")
)

type PvzService interface {
	DummyLogin(ctx context.Context, role string) (*string, werrors.Werror)
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

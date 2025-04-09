package service

import (
	"log/slog"
	"pvz-service/internal/app/db"
	"pvz-service/internal/tokenizer"
	"pvz-service/pkg/cryptor"
)

type PvzService interface {
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

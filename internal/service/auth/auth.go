package auth

import (
	"context"
	"log/slog"
	"net/http"
	"pvz-service/internal/apperrors"
	"pvz-service/internal/db"
	"pvz-service/internal/enum"
	"pvz-service/internal/tokenizer"
	"pvz-service/pkg/cryptor"
	"pvz-service/pkg/werrors"
)

type AuthService interface {
	DummyLogin(ctx context.Context, role string) (*string, werrors.Werror)
}

type authService struct {
	storage   db.DB
	logger    *slog.Logger
	cryptor   cryptor.Cryptor
	tokenizer tokenizer.Tokenizer
}

func New(storage db.DB, logger *slog.Logger, cryptor cryptor.Cryptor, tok tokenizer.Tokenizer) AuthService {
	return &authService{
		storage:   storage,
		logger:    logger,
		cryptor:   cryptor,
		tokenizer: tok,
	}
}

func (s *authService) DummyLogin(ctx context.Context, role string) (*string, werrors.Werror) {
	if !enum.CheckRole(role) {
		return nil, werrors.New(apperrors.ErrInvalidRole, http.StatusBadRequest)
	}

	token, err := s.tokenizer.GenerateToken(role)
	if err != nil {
		s.logger.Error("generate token: " + err.Error())
		return nil, werrors.New(apperrors.ErrSmthWentWrong, http.StatusInternalServerError)
	}

	return token, nil
}

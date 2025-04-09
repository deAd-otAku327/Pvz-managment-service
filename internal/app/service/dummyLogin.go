package service

import (
	"context"
	"net/http"
	"pvz-service/internal/models"
	"pvz-service/pkg/werrors"
)

func (s *pvzService) DummyLogin(ctx context.Context, role string) (*string, werrors.Werror) {
	if !models.CheckRole(role) {
		return nil, werrors.New(errInvalidRole, http.StatusBadRequest)
	}

	token, err := s.tokenizer.GenerateToken(role)
	if err != nil {
		s.logger.Error("generate token: " + err.Error())
		return nil, werrors.New(errSmthWentWrong, http.StatusInternalServerError)
	}

	return token, nil
}

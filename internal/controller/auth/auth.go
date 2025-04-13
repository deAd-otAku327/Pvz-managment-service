package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pvz-service/internal/apperrors"
	"pvz-service/internal/middleware"
	"pvz-service/internal/service"
	"pvz-service/pkg/response"
)

type AuthController interface {
	DummyLogin() http.HandlerFunc
	Register() http.HandlerFunc
	Login() http.HandlerFunc
}

type authController struct {
	service service.Service
}

func New(service service.Service) AuthController {
	return &authController{
		service: service,
	}
}

// No DTO for dummy.
func (c *authController) DummyLogin() http.HandlerFunc {
	type dummyLoginRequest struct {
		Role string `json:"role"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		request := dummyLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil || request.Role == "" {
			response.MakeErrorResponseJSON(w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody)
			return
		}

		token, serviceErr := c.service.DummyLogin(r.Context(), request.Role)
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		w.Header().Set("Set-Cookie", fmt.Sprintf("%s=%s", middleware.CookieName, *token))

		response.MakeResponseJSON(w, http.StatusOK, nil)
	}
}

func (c *authController) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.MakeResponseJSON(w, http.StatusNotImplemented, nil)
	}
}

func (c *authController) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.MakeResponseJSON(w, http.StatusNotImplemented, nil)
	}
}

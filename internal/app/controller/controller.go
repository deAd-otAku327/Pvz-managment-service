package controller

import (
	"errors"
	"net/http"
	"pvz-service/internal/app/service"
)

var errInvalidRequestBody = errors.New("invalid request body provided")

type Controller interface {
	DummyLogin() http.HandlerFunc

	Register() http.HandlerFunc
	Login() http.HandlerFunc

	CreatePvz() http.HandlerFunc
	GetPvzList() http.HandlerFunc
	CloseLastReception() http.HandlerFunc
	DeleteLastProduct() http.HandlerFunc
	CreateReception() http.HandlerFunc
	AddProduct() http.HandlerFunc
}

type controller struct {
	service service.PvzService
}

func New(s service.PvzService) Controller {
	return &controller{
		service: s,
	}
}

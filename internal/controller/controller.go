package controller

import (
	"net/http"
	"pvz-service/internal/service"
)

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
	service service.Service
}

func New(s service.Service) Controller {
	return &controller{
		service: s,
	}
}

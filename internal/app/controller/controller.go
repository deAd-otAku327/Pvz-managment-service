package controller

import (
	"net/http"
	"pvz-service/internal/app/service"
)

type Controller interface {
	DummyLogin() http.HandlerFunc
}

type controller struct {
	service service.PvzService
}

func New(s service.PvzService) Controller {
	return &controller{
		service: s,
	}
}

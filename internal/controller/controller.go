package controller

import (
	"net/http"
	"pvz-service/internal/controller/auth"
	"pvz-service/internal/controller/product"
	"pvz-service/internal/controller/pvz"
	"pvz-service/internal/controller/reception"
	"pvz-service/internal/service"
)

type Controller interface {
	auth.AuthController
	pvz.PvzController
	reception.ReceptionController
	product.ProductController
}

type controller struct {
	authController      auth.AuthController
	pvzController       pvz.PvzController
	receptionController reception.ReceptionController
	productController   product.ProductController
}

func New(s service.Service) Controller {
	return &controller{
		authController:      auth.New(s),
		pvzController:       pvz.New(s),
		receptionController: reception.New(s),
		productController:   product.New(s),
	}
}

func (c *controller) DummyLogin() http.HandlerFunc {
	return c.authController.DummyLogin()
}

func (c *controller) Register() http.HandlerFunc {
	return c.authController.Register()
}

func (c *controller) Login() http.HandlerFunc {
	return c.authController.Login()
}

func (c *controller) CreatePvz() http.HandlerFunc {
	return c.pvzController.CreatePvz()
}

func (c *controller) GetPvzList() http.HandlerFunc {
	return c.pvzController.GetPvzList()
}

func (c *controller) AddProduct() http.HandlerFunc {
	return c.productController.AddProduct()
}

func (c *controller) DeleteProduct() http.HandlerFunc {
	return c.productController.DeleteProduct()
}

func (c *controller) CreateReception() http.HandlerFunc {
	return c.receptionController.CreateReception()
}

func (c *controller) CloseReception() http.HandlerFunc {
	return c.receptionController.CloseReception()
}

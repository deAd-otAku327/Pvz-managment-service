package product

import (
	"encoding/json"
	"net/http"
	"pvz-service/internal/apperrors"
	"pvz-service/internal/controller/shared/consts"
	"pvz-service/internal/dto"
	"pvz-service/internal/service"
	"pvz-service/pkg/response"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductController interface {
	AddProduct() http.HandlerFunc
	DeleteProduct() http.HandlerFunc
}

type productController struct {
	service service.Service
}

func New(service service.Service) ProductController {
	return &productController{
		service: service,
	}
}

func (c *productController) AddProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.AddProductRequestDTO{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			response.MakeErrorResponseJSON(w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody)
			return
		}

		product, serviceErr := c.service.AddProduct(r.Context(), &request)
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		response.MakeResponseJSON(w, http.StatusCreated, product)
	}
}

func (c *productController) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Regexps on routes guarantees no error.
		p, _ := strconv.Atoi(mux.Vars(r)[consts.URLParamPvzID]) //nolint:errcheck
		request := dto.DeleteProductRequestDTO{
			PvzID: p,
		}

		serviceErr := c.service.DeleteProduct(r.Context(), &request)
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		response.MakeResponseJSON(w, http.StatusOK, nil)
	}
}

package controller

import (
	"encoding/json"
	"net/http"
	"pvz-service/internal/apperrors"
	"pvz-service/internal/dto"
	"pvz-service/pkg/response"
)

func (c *controller) AddProduct() http.HandlerFunc {
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

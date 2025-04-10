package controller

import (
	"encoding/json"
	"net/http"
	"pvz-service/pkg/response"
)

func (c *controller) AddProduct() http.HandlerFunc {
	type addProductRequest struct {
		Type  string `yaml:"type"`
		PvzID int    `yaml:"pvzId"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		request := addProductRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil || request.Type == "" || request.PvzID == 0 {
			response.MakeErrorResponseJSON(w, http.StatusBadRequest, errInvalidRequestBody)
			return
		}

		product, serviceErr := c.service.AddProduct(r.Context(), request.Type, request.PvzID)
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		response.MakeResponseJSON(w, http.StatusCreated, product)
	}
}

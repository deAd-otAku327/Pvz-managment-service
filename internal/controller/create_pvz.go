package controller

import (
	"encoding/json"
	"net/http"
	"pvz-service/pkg/response"
)

func (c *controller) CreatePvz() http.HandlerFunc {
	type createPvzRequest struct {
		City string `json:"city"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		request := createPvzRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil || request.City == "" {
			response.MakeErrorResponseJSON(w, http.StatusBadRequest, errInvalidRequestBody)
			return
		}

		pvz, serviceErr := c.service.CreatePvz(r.Context(), request.City)
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		response.MakeResponseJSON(w, http.StatusCreated, pvz)
	}
}

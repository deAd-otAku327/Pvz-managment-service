package controller

import (
	"encoding/json"
	"net/http"
	"pvz-service/internal/apperrors"
	"pvz-service/internal/dto"
	"pvz-service/pkg/response"
)

func (c *controller) CreateReception() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.CreateReceptionRequestDTO{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			response.MakeErrorResponseJSON(w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody)
			return
		}

		reception, serviceErr := c.service.CreateReception(r.Context(), &request)
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		response.MakeResponseJSON(w, http.StatusCreated, reception)
	}
}

package controller

import (
	"encoding/json"
	"net/http"
	"pvz-service/pkg/response"
)

func (c *controller) CreateReception() http.HandlerFunc {
	type createReceptionRequest struct {
		PvzID int `yaml:"pvzId"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		request := createReceptionRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil || request.PvzID == 0 {
			response.MakeErrorResponseJSON(w, http.StatusBadRequest, errInvalidRequestBody)
			return
		}

		reception, serviceErr := c.service.CreateReception(r.Context(), request.PvzID)
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		response.MakeResponseJSON(w, http.StatusOK, reception)
	}
}

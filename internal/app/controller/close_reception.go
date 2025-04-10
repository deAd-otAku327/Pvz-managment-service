package controller

import (
	"net/http"
	"pvz-service/pkg/response"

	"github.com/gorilla/mux"
)

const paramPvzID = "pvzId"

func (c *controller) CloseLastReception() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pvzId := mux.Vars(r)[paramPvzID]

		reception, serviceErr := c.service.CloseReception(r.Context(), pvzId)
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		response.MakeResponseJSON(w, http.StatusCreated, reception)
	}
}

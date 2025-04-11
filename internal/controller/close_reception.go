package controller

import (
	"net/http"
	"pvz-service/internal/dto"
	"pvz-service/pkg/response"
	"strconv"

	"github.com/gorilla/mux"
)

const paramPvzID = "pvzId"

func (c *controller) CloseLastReception() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Regexps on routes guarantees no error.
		p, _ := strconv.Atoi(mux.Vars(r)[paramPvzID]) //nolint:errcheck
		request := dto.CloseReceptionRequestDTO{
			PvzID: p,
		}

		reception, serviceErr := c.service.CloseReception(r.Context(), &request)
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		response.MakeResponseJSON(w, http.StatusOK, reception)
	}
}

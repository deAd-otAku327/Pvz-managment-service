package controller

import (
	"net/http"
	"pvz-service/internal/dto"
	"pvz-service/pkg/response"
	"strconv"

	"github.com/gorilla/mux"
)

func (c *controller) DeleteLastProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Regexps on routes guarantees no error.
		p, _ := strconv.Atoi(mux.Vars(r)[paramPvzID]) //nolint:errcheck
		request := dto.DeleteProductRequestDTO{
			PvzId: p,
		}

		serviceErr := c.service.DeleteProduct(r.Context(), &request)
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		response.MakeResponseJSON(w, http.StatusOK, nil)
	}
}

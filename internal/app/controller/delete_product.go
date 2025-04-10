package controller

import (
	"net/http"
	"pvz-service/pkg/response"

	"github.com/gorilla/mux"
)

func (c *controller) DeleteLastProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pvzId := mux.Vars(r)[paramPvzID]

		serviceErr := c.service.DeleteProduct(r.Context(), pvzId)
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		response.MakeResponseJSON(w, http.StatusOK, nil)
	}
}

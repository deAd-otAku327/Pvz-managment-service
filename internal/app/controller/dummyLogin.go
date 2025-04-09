package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pvz-service/pkg/response"
)

func (c *controller) DummyLogin() http.HandlerFunc {
	type dummyLoginRequest struct {
		Role string `json:"role"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		request := dummyLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil || request.Role == "" {
			response.MakeErrorResponseJSON(w, http.StatusBadRequest, errInvalidRequestBody)
			return
		}

		token, serviceErr := c.service.DummyLogin(r.Context(), request.Role)
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		w.Header().Set("Set-Cookie", fmt.Sprintf("token=%s", *token))

		response.MakeResponseJSON(w, http.StatusOK, nil)
	}
}

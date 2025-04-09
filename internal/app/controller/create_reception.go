package controller

import "net/http"

func (c *controller) CreateReception() http.HandlerFunc {
	type createReceptionRequest struct {
	}
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

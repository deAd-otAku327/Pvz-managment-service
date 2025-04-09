package controller

import "net/http"

func (c *controller) Register() http.HandlerFunc {
	type registerRequest struct {
	}
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

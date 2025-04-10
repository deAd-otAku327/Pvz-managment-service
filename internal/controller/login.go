package controller

import "net/http"

func (c *controller) Login() http.HandlerFunc {
	type loginRequest struct {
	}
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

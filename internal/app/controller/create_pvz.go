package controller

import "net/http"

func (c *controller) CreatePvz() http.HandlerFunc {
	type createPvzRequest struct {
	}
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

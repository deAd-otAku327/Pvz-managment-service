package controller

import "net/http"

func (c *controller) DeleteLastProduct() http.HandlerFunc {
	type deleteLastProductRequest struct {
	}
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

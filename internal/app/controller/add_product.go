package controller

import "net/http"

func (c *controller) AddProduct() http.HandlerFunc {
	type addProductRequest struct {
	}
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

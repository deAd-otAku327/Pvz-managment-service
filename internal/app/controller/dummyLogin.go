package controller

import (
	"net/http"
)

func (c *controller) DummyLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

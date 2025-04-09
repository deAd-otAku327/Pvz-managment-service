package controller

import "net/http"

func (c *controller) GetPvzList() http.HandlerFunc {
	type getPvzListRequest struct {
	}
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

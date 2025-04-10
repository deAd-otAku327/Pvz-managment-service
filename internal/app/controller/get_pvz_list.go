package controller

import (
	"net/http"
	"pvz-service/pkg/response"
)

const (
	paramStartDate = "startDate"
	paramEndDate   = "endDate"
	paramPage      = "page"
	paramLimit     = "limit"
)

func (c *controller) GetPvzList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startDate := r.URL.Query().Get(paramStartDate)
		endDate := r.URL.Query().Get(paramEndDate)
		page := r.URL.Query().Get(paramPage)
		limit := r.URL.Query().Get(paramLimit)

		summaryInfo, serviceErr := c.service.GetPvzList(r.Context(), startDate, endDate, page, limit)
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		response.MakeResponseJSON(w, http.StatusOK, summaryInfo)
	}
}

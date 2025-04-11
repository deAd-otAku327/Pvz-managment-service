package controller

import (
	"net/http"
	"pvz-service/internal/apperrors"
	"pvz-service/internal/dto"
	"pvz-service/pkg/response"

	"github.com/gorilla/schema"
)

const (
	paramStartDate = "startDate"
	paramEndDate   = "endDate"
	paramPage      = "page"
	paramLimit     = "limit"
)

func (c *controller) GetPvzList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			response.MakeErrorResponseJSON(w, http.StatusBadRequest, apperrors.ErrInvalidRequestParams)
			return
		}

		request := dto.PvzFilterParamsDTO{}
		err = schema.NewDecoder().Decode(&request, r.Form)
		if err != nil {
			response.MakeErrorResponseJSON(w, http.StatusBadRequest, apperrors.ErrInvalidRequestParams)
			return
		}

		summaryInfo, serviceErr := c.service.GetPvzList(r.Context(), &request)
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		response.MakeResponseJSON(w, http.StatusOK, summaryInfo)
	}
}

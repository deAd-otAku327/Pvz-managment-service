package reception

import (
	"encoding/json"
	"net/http"
	"pvz-service/internal/apperrors"
	"pvz-service/internal/controller/shared/consts"
	"pvz-service/internal/dto"
	"pvz-service/internal/service"
	"pvz-service/pkg/response"
	"strconv"

	"github.com/gorilla/mux"
)

type ReceptionController interface {
	CreateReception() http.HandlerFunc
	CloseReception() http.HandlerFunc
}

type receptionController struct {
	service service.Service
}

func New(service service.Service) ReceptionController {
	return &receptionController{
		service: service,
	}
}

func (c *receptionController) CreateReception() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.CreateReceptionRequestDTO{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			response.MakeErrorResponseJSON(w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody)
			return
		}

		reception, serviceErr := c.service.CreateReception(r.Context(), &request)
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		response.MakeResponseJSON(w, http.StatusCreated, reception)
	}
}

func (c *receptionController) CloseReception() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Regexps on routes guarantees no error.
		p, _ := strconv.Atoi(mux.Vars(r)[consts.URLParamPvzID]) //nolint:errcheck
		request := dto.CloseReceptionRequestDTO{
			PvzID: p,
		}

		reception, serviceErr := c.service.CloseReception(r.Context(), &request)
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		response.MakeResponseJSON(w, http.StatusOK, reception)
	}
}

package pvz

import (
	"encoding/json"
	"net/http"
	"pvz-service/internal/apperrors"
	"pvz-service/internal/dto"
	dtomap "pvz-service/internal/mappers/dto"
	"pvz-service/internal/service"
	"pvz-service/pkg/response"

	"github.com/gorilla/schema"
)

type PvzController interface {
	CreatePvz() http.HandlerFunc
	GetPvzList() http.HandlerFunc
}

type pvzController struct {
	service service.Service
}

func New(service service.Service) PvzController {
	return &pvzController{
		service: service,
	}
}

func (c *pvzController) CreatePvz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.CreatePvzRequestDTO{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			response.MakeErrorResponseJSON(w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody)
			return
		}

		pvz, serviceErr := c.service.CreatePvz(r.Context(), dtomap.MapToPvzCreate(&request))
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		response.MakeResponseJSON(w, http.StatusCreated, pvz)
	}
}

func (c *pvzController) GetPvzList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			response.MakeErrorResponseJSON(w, http.StatusBadRequest, apperrors.ErrInvalidRequestParams)
			return
		}

		request := dto.PvzFilterParamsDTO{}
		err = schema.NewDecoder().Decode(&request, r.Form)
		if err != nil {
			response.MakeErrorResponseJSON(w, http.StatusBadRequest, apperrors.ErrInvalidParamFormat)
			return
		}

		summaryInfo, serviceErr := c.service.GetPvzList(r.Context(), dtomap.MapToPvzFilterParams(&request))
		if serviceErr != nil {
			response.MakeErrorResponseJSON(w, serviceErr.Code(), serviceErr)
			return
		}

		response.MakeResponseJSON(w, http.StatusOK, summaryInfo)
	}
}

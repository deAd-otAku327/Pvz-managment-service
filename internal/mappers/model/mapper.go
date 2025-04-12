package model

import (
	"pvz-service/internal/dto"
	"pvz-service/internal/models"
)

func MapToPvzResponse(pvz *models.Pvz) *dto.PvzResponseDTO {
	return &dto.PvzResponseDTO{
		ID:               pvz.ID,
		RegistrationDate: pvz.RegistrationDate.Format("2006-01-02 15:04:05"),
		City:             pvz.City,
	}
}

func MapToGetPvzListResponse(pvzList *models.PvzList) *dto.GetPvzListResponseDTO {
	return &dto.GetPvzListResponseDTO{}
}

func MapToReceptionResponse(reception *models.Reception) *dto.ReceptionResponseDTO {
	return &dto.ReceptionResponseDTO{}
}

func MapToProductResponse(product *models.Product) *dto.ProductResponseDTO {
	return &dto.ProductResponseDTO{}
}

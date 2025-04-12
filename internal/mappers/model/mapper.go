package model

import (
	"pvz-service/internal/dto"
	"pvz-service/internal/models"
)

func MapToPvzResponse(pvz *models.Pvz) *dto.PvzResponseDTO {
	return &dto.PvzResponseDTO{}
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

package model

import (
	"pvz-service/internal/dto"
	"pvz-service/internal/models"
)

const DateTimeFormat = "2006-01-02 15:04:05"

func MapToPvzResponse(pvz *models.Pvz) *dto.PvzResponseDTO {
	return &dto.PvzResponseDTO{
		ID:               pvz.ID,
		RegistrationDate: pvz.RegistrationDate.Format(DateTimeFormat),
		City:             pvz.City,
	}
}

func MapToGetPvzListResponse(pvzList *models.PvzList) *dto.GetPvzListResponseDTO {
	return &dto.GetPvzListResponseDTO{}
}

func MapToReceptionResponse(reception *models.Reception) *dto.ReceptionResponseDTO {
	return &dto.ReceptionResponseDTO{
		ID:       reception.ID,
		DateTime: reception.DateTime.Format(DateTimeFormat),
		PvzID:    reception.PvzID,
		Status:   reception.Status,
	}
}

func MapToProductResponse(product *models.Product) *dto.ProductResponseDTO {
	return &dto.ProductResponseDTO{
		ID:          product.ID,
		DateTime:    product.DateTime,
		ReceptionID: product.ReceptionID,
		Type:        product.Type,
	}
}

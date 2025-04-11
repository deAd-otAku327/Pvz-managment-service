package entity

import (
	"pvz-service/internal/dto"
	"pvz-service/internal/entities"
)

func MapToPvzResponse(pvz *entities.Pvz) *dto.PvzResponseDTO {
	return &dto.PvzResponseDTO{}
}

func MapToGetPvzListResponse(pvzs []entities.Pvz, receptions []entities.Reception) *dto.GetPvzListResponseDTO {
	return &dto.GetPvzListResponseDTO{}
}

func MapToProductResponse(p *entities.Product) *dto.ProductResponseDTO {
	return &dto.ProductResponseDTO{}
}

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

func MapToGetPvzListResponse(pvzList models.PvzList) *dto.GetPvzListResponseDTO {
	pvzListResponse := make(dto.GetPvzListResponseDTO, 0)
	for _, pwr := range pvzList {
		pvzWithReceptions := dto.PvzWithReceptionsDTO{
			Pvz: MapToPvzResponse(pwr.Pvz),
		}
		for _, r := range pwr.Receptions {
			receptionWithProducts := dto.ReceptionWithProductsDTO{
				Reception: MapToReceptionResponse(r),
			}
			for _, p := range r.Products {
				receptionWithProducts.Products = append(receptionWithProducts.Products, MapToProductResponse(p))
			}
			pvzWithReceptions.Receptions = append(pvzWithReceptions.Receptions, &receptionWithProducts)
		}
		pvzListResponse = append(pvzListResponse, &pvzWithReceptions)
	}

	return &pvzListResponse
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
		DateTime:    product.DateTime.Format(DateTimeFormat),
		ReceptionID: product.ReceptionID,
		Type:        product.Type,
	}
}

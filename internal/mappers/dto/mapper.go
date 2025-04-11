package dto

import (
	"pvz-service/internal/dto"
	"pvz-service/internal/models"
)

func MapToPvzFilterParams(fp *dto.PvzFilterParamsDTO) *models.PvzFilterParams {
	return &models.PvzFilterParams{
		StartDate: fp.StartDate.Date,
		EndDate:   fp.EndDate.Date,
		Page:      fp.Page,
		Limit:     fp.Limit,
	}
}

func MapToPvzCreate(cpr *dto.CreatePvzRequestDTO) *models.PvzCreate {
	return &models.PvzCreate{
		City: cpr.City,
	}
}

func MapToAddProduct(apr *dto.AddProductRequestDTO) *models.AddProduct {
	return &models.AddProduct{
		Type: apr.Type,
	}
}

func MapToDeleteProduct(dpr *dto.DeleteProductRequestDTO) *models.DeleteProduct {
	return &models.DeleteProduct{
		PvzID: dpr.PvzID,
	}
}

func MapToCreateReception(crr *dto.CreateReceptionRequestDTO) *models.CreateReception {
	return &models.CreateReception{
		PvzID: crr.PvzID,
	}
}

func MapToCloseReception(crr *dto.CloseReceptionRequestDTO) *models.CloseReception {
	return &models.CloseReception{
		PvzID: crr.PvzID,
	}
}

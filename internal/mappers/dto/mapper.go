package dto

import (
	"pvz-service/internal/dto"
	"pvz-service/internal/models"
	"time"
)

func MapToPvzFilterParams(fp *dto.PvzFilterParamsDTO) *models.PvzFilterParams {
	endDate := fp.EndDate.Date
	if endDate.Equal(time.Time{}) {
		endDate = time.Now()
	}
	return &models.PvzFilterParams{
		StartDate: fp.StartDate.Date,
		EndDate:   endDate,
		Page:      fp.Page,
		Limit:     fp.Limit,
	}
}

func MapToPvzCreate(cpr *dto.CreatePvzRequestDTO) *models.CreatePvz {
	return &models.CreatePvz{
		City: cpr.City,
	}
}

func MapToAddProduct(apr *dto.AddProductRequestDTO) *models.AddProduct {
	return &models.AddProduct{
		PvzID: apr.PvzID,
		Type:  apr.Type,
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

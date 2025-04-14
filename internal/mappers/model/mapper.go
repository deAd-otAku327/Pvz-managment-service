package model

import (
	"pvz-service/internal/dto"
	"pvz-service/internal/entities"
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

func MapToAddProduct(addProduct *models.AddProduct) *entities.AddProduct {
	return &entities.AddProduct{
		Type:  addProduct.Type,
		PvzID: addProduct.PvzID,
	}
}

func MapToDeleteProduct(deleteProduct *models.DeleteProduct) *entities.DeleteProduct {
	return &entities.DeleteProduct{
		PvzID: deleteProduct.PvzID,
	}
}

func MapToCreatePvz(createPvz *models.CreatePvz) *entities.CreatePvz {
	return &entities.CreatePvz{
		City: createPvz.City,
	}
}

func MapToPvzFilterParams(filters *models.PvzFilterParams) *entities.PvzFilterParams {
	return &entities.PvzFilterParams{
		StartDate: filters.StartDate,
		EndDate:   filters.EndDate,
		Page:      filters.Page,
		Limit:     filters.Limit,
	}
}
func MapToCreateReception(createReception *models.CreateReception) *entities.CreateReception {
	return &entities.CreateReception{
		PvzID: createReception.PvzID,
	}
}

func MapToCloseReception(closeReception *models.CloseReception) *entities.CloseReception {
	return &entities.CloseReception{
		PvzID: closeReception.PvzID,
	}
}

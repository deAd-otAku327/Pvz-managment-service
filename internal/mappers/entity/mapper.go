package entity

import (
	"pvz-service/internal/entities"
	"pvz-service/internal/models"
)

func MapToPvzList(pvzs []entities.Pvz, recepts []entities.Reception, products []entities.Product) *models.PvzList {
	return &models.PvzList{}
}

func MapToPvz(pvz entities.Pvz) *models.Pvz {
	return &models.Pvz{}
}

func MapToReception(reception *entities.Reception) *models.Reception {
	return &models.Reception{}
}

func MapToProduct(product *entities.Product) *models.Product {
	return &models.Product{}
}

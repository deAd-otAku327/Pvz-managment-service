package entity

import (
	"pvz-service/internal/entities"
	"pvz-service/internal/models"
)

func MapToPvzList(pvzs []*entities.Pvz, receptions []*entities.Reception, products []*entities.Product) models.PvzList {
	receptionIDToProducts := make(map[int][]*entities.Product)
	for _, p := range products {
		receptionIDToProducts[p.ReceptionID] = append(receptionIDToProducts[p.ReceptionID], p)
	}

	pvzList := make(models.PvzList, 0)
	for _, pvz := range pvzs {
		pvzWithReceptions := models.PvzWithReceptions{
			Pvz: MapToPvz(pvz),
		}
		for _, r := range receptions {
			if pvz.ID == r.PvzID {
				pvzWithReceptions.Receptions = append(pvzWithReceptions.Receptions, MapToReception(r, receptionIDToProducts[r.ID]))
			}
		}
		pvzList = append(pvzList, &pvzWithReceptions)
	}

	return pvzList
}

func MapToPvz(pvz *entities.Pvz) *models.Pvz {
	return &models.Pvz{
		ID:               pvz.ID,
		RegistrationDate: pvz.RegistrationDate,
		City:             pvz.City,
	}
}

func MapToReception(reception *entities.Reception, prods []*entities.Product) *models.Reception {
	return &models.Reception{
		ID:       reception.ID,
		DateTime: reception.DateTime,
		PvzID:    reception.PvzID,
		Products: func() []*models.Product {
			if prods != nil {
				products := make([]*models.Product, 0)
				for _, p := range prods {
					products = append(products, MapToProduct(p))
				}
				if len(products) > 0 {
					return products
				}
			}

			return nil
		}(),
		Status: reception.Status,
	}
}

func MapToProduct(product *entities.Product) *models.Product {
	return &models.Product{
		ID:          product.ID,
		DateTime:    product.DateTime,
		ReceptionID: product.ReceptionID,
		Type:        product.Type,
	}
}

package models

import (
	"pvz-service/internal/apperrors"
	"pvz-service/internal/enum"
)

type AddProduct struct {
	Type  string
	PvzID int
}

type DeleteProduct struct {
	PvzID int
}

func (ap *AddProduct) Validate() error {
	if !enum.CheckProductType(ap.Type) {
		return apperrors.ErrInvalidProductType
	}

	if ap.PvzID <= 0 {
		return apperrors.ErrInvalidPvzID
	}

	return nil
}

func (dp *DeleteProduct) Validate() error {
	if dp.PvzID <= 0 {
		return apperrors.ErrInvalidPvzID
	}

	return nil
}

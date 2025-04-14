package models

import (
	"pvz-service/internal/apperrors"
	"pvz-service/internal/enum"
	"time"
)

type AddProduct struct {
	Type  string
	PvzID int
}

type DeleteProduct struct {
	PvzID int
}

type Product struct {
	ID          int
	DateTime    time.Time
	ReceptionID int
	Type        string
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

func (p *Product) Validate() error {
	if !enum.CheckProductType(p.Type) {
		return apperrors.ErrInvalidProductType
	}

	if p.ReceptionID <= 0 {
		return apperrors.ErrInvalidReceptionID
	}

	return nil
}

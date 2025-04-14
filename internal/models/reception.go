package models

import (
	"pvz-service/internal/apperrors"
	"pvz-service/internal/enum"
	"time"
)

type CreateReception struct {
	PvzID int
}

type CloseReception struct {
	PvzID int
}

type Reception struct {
	ID       int
	DateTime time.Time
	PvzID    int
	Products []*Product
	Status   string
}

func (cr *CreateReception) Validate() error {
	if cr.PvzID <= 0 {
		return apperrors.ErrInvalidPvzID
	}

	return nil
}

func (cr *CloseReception) Validate() error {
	if cr.PvzID <= 0 {
		return apperrors.ErrInvalidPvzID
	}

	return nil
}

func (r *Reception) Validate() error {
	if !enum.CheckStatus(r.Status) {
		return apperrors.ErrInvalidReceptionStatus
	}

	if r.ID <= 0 {
		return apperrors.ErrInvalidReceptionID
	}

	if r.PvzID <= 0 {
		return apperrors.ErrInvalidPvzID
	}

	for _, r := range r.Products {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	return nil
}

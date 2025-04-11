package models

import "pvz-service/internal/apperrors"

type CreateReception struct {
	PvzID int
}

type CloseReception struct {
	PvzID int
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

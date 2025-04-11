package models

import (
	"pvz-service/internal/apperrors"
	"pvz-service/internal/enum"
	"time"
)

const (
	minPvzFilterPage = 1

	minPvzFilterLimit = 1
	maxPvzFilterLimit = 30
)

type PvzCreate struct {
	City string
}

type PvzFilterParams struct {
	StartDate time.Time
	EndDate   time.Time
	Page      int
	Limit     int
}

func (pc *PvzCreate) Validate() error {
	if !enum.CheckCity(pc.City) {
		return apperrors.ErrInvalidCity
	}

	return nil
}

func (fp *PvzFilterParams) Validate() error {
	if fp.Page < minPvzFilterPage {
		return apperrors.ErrInvalidPageParam
	}

	if fp.Limit < minPvzFilterLimit || fp.Limit > maxPvzFilterLimit {
		return apperrors.ErrInvalidLimitParam
	}

	return nil
}

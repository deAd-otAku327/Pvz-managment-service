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

type PvzList []*PvzWithReceptions

type PvzWithReceptions struct {
	Pvz        *Pvz
	Receptions []*ReceptionWithProducts
}

type Pvz struct {
	ID               int
	RegistrationDate time.Time
	City             string
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

func (pl PvzList) Validate() error {
	return func() error {
		for _, r := range pl {
			if err := r.Validate(); err != nil {
				return err
			}
		}

		return nil
	}()
}

func (pwr *PvzWithReceptions) Validate() error {
	return func() error {
		if err := pwr.Pvz.Validate(); err != nil {
			return err
		}

		for _, r := range pwr.Receptions {
			if err := r.Validate(); err != nil {
				return err
			}
		}

		return nil
	}()
}

func (p *Pvz) Validate() error {
	if !enum.CheckCity(p.City) {
		return apperrors.ErrInvalidCity
	}

	if p.ID <= 0 {
		return apperrors.ErrInvalidPvzID
	}

	return nil
}

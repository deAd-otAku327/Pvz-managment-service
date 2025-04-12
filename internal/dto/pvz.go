package dto

import (
	"time"
)

type CreatePvzRequestDTO struct {
	City string `json:"city"`
}

type GetPvzListResponseDTO []*PvzWithReceptionsDTO

type PvzWithReceptionsDTO struct {
	Pvz        *PvzResponseDTO             `json:"pvz"`
	Receptions []*ReceptionWithProductsDTO `json:"receptions"`
}

type PvzResponseDTO struct {
	ID               int    `json:"id"`
	RegistrationDate string `json:"registrationDate"`
	City             string `json:"city"`
}

// Page value = 0 will be recognized as default value = 1.
type PvzFilterParamsDTO struct {
	StartDate DateParam `schema:"startDate"`
	EndDate   DateParam `schema:"endDate"`
	Page      int       `schema:"page,default:1"`
	Limit     int       `schema:"limit,default:10"`
}

type DateParam struct {
	Date time.Time
}

// Implementation of gorilla/schema interface.
func (dp *DateParam) UnmarshalText(text []byte) error {
	parsedTime, err := time.Parse("2006-01-02", string(text))
	if err != nil {
		return err
	}
	dp.Date = parsedTime
	return nil
}

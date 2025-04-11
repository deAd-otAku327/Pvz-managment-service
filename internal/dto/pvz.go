package dto

import "time"

type CreatePvzRequestDTO struct {
	City string `json:"city"`
}

type GetPvzListResponseDTO []PvzWithReceptionsDTO

type PvzWithReceptionsDTO struct {
	Pvz        PvzResponseDTO             `json:"pvz"`
	Receptions []ReceptionWithProductsDTO `json:"receptions"`
}

type PvzResponseDTO struct {
	ID               string    `json:"id,omitempty"`
	RegistrationDate time.Time `json:"registrationDate"`
	City             string    `json:"city"`
}

type PvzFilterParamsDTO struct {
	StartDate time.Time `schema:"startDate"`
	EndDate   time.Time `schema:"endDate"`
	Page      int       `schema:"page,omitempty"`
	Limit     int       `schema:"limit,omitempty"`
}

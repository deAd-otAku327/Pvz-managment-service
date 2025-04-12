package dto

import "time"

type CreateReceptionRequestDTO struct {
	PvzID int `json:"pvzId"`
}

type CloseReceptionRequestDTO struct {
	PvzID int
}

type ReceptionWithProductsDTO struct {
	Reception ReceptionResponseDTO `json:"reception"`
	Products  []ProductResponseDTO `json:"products"`
}

type ReceptionResponseDTO struct {
	ID       int       `json:"id"`
	DateTime time.Time `json:"dateTime"`
	PvzID    int       `json:"pvzId"`
	Status   string    `json:"status"`
}

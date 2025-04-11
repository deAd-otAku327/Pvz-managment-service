package dto

import "time"

type AddProductRequestDTO struct {
	Type  string `json:"type"`
	PvzID int    `json:"pvzId"`
}

type DeleteProductRequestDTO struct {
	PvzID int
}

type ProductResponseDTO struct {
	ID          int       `json:"id,omitempty"`
	DateTime    time.Time `json:"dateTime"`
	Type        string    `json:"type"`
	ReceptionID string    `json:"receptionId"`
}

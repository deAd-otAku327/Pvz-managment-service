package dto

type AddProductRequestDTO struct {
	Type  string `json:"type"`
	PvzID int    `json:"pvzId"`
}

type DeleteProductRequestDTO struct {
	PvzID int
}

type ProductResponseDTO struct {
	ID          int    `json:"id"`
	DateTime    string `json:"dateTime"`
	Type        string `json:"type"`
	ReceptionID int    `json:"receptionId"`
}

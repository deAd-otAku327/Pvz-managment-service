package dto

type CreateReceptionRequestDTO struct {
	PvzID int `json:"pvzId"`
}

type CloseReceptionRequestDTO struct {
	PvzID int
}

type ReceptionWithProductsDTO struct {
	Reception *ReceptionResponseDTO `json:"reception"`
	Products  []*ProductResponseDTO `json:"products"`
}

type ReceptionResponseDTO struct {
	ID       int    `json:"id"`
	DateTime string `json:"dateTime"`
	PvzID    int    `json:"pvzId"`
	Status   string `json:"status"`
}

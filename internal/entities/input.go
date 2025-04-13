package entities

import "time"

type AddProduct struct {
	Type  string
	PvzID int
}

type DeleteProduct struct {
	PvzID int
}

type CreateReception struct {
	PvzID int
}

type CloseReception struct {
	PvzID int
}

type CreatePvz struct {
	City string
}

type PvzFilterParams struct {
	StartDate time.Time
	EndDate   time.Time
	Page      int
	Limit     int
}

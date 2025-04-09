package models

import "time"

type User struct {
	ID      int
	Email   string `json:"email"`
	RawRole UserRole
	Role    string `json:"role"`
}

type PVZ struct {
	ID               int
	RegistrationDate time.Time `json:"registrationDate"`
	RawCity          City
	City             string `json:"city"`
}

type Reception struct {
	ID        int
	DateTime  time.Time `json:"dateTime"`
	PvzID     int       `json:"pvzId"`
	RawStatus ReceptionStatus
	Status    string `json:"status"`
}

type Product struct {
	ID          int
	DateTime    time.Time
	RawType     ProductType
	Type        string `json:"type"`
	ReceptionID int    `json:"receptionId"`
}

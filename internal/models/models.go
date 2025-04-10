package models

import "time"

type User struct {
	ID    int
	Email string `json:"email"`
	Role  string `json:"role"`
}

type PVZ struct {
	ID               int       `json:"id"`
	RegistrationDate time.Time `json:"registrationDate"`
	City             string    `json:"city"`
}

type Reception struct {
	ID       int       `json:"id"`
	DateTime time.Time `json:"dateTime"`
	PvzID    int       `json:"pvzId"`
	Status   string    `json:"status"`
}

type Product struct {
	ID          int
	DateTime    time.Time `json:"dateTime"`
	Type        string    `json:"type"`
	ReceptionID int       `json:"receptionId"`
}

type SummaryInfo []PVZSummary

type PVZSummary struct {
	Pvz        PVZ                `json:"pvz"`
	Receptions []ReceptionSummary `json:"receptions"`
}

type ReceptionSummary struct {
	Reception Reception `json:"reception"`
	Products  []Product `json:"product"`
}

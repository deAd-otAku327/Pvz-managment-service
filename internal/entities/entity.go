package entities

import "time"

type User struct {
	ID    int
	Email string
	Role  string
}

type Pvz struct {
	ID               int
	RegistrationDate time.Time
	City             string
}

type Reception struct {
	ID       int
	DateTime time.Time
	PvzID    int
	Products []Product
	Status   string
}

type Product struct {
	ID       int
	DateTime time.Time
	Type     string
}

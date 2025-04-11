package models

import "time"

const (
	defaultPage = 1
	minPage     = 1

	defaultLimit = 10
	minLimit     = 1
	maxLimit     = 30

	datetimeLayout = "2006-01-02 15:04"
)

type PvzCreate struct {
	City string
}

type PvzFilterParams struct {
	StartDate time.Time
	EndDate   time.Time
	Page      int
	Limit     int
}

func (fp *PvzFilterParams) Validate() error {
	return nil
}

func (pc *PvzCreate) Validate() error {
	return nil
}

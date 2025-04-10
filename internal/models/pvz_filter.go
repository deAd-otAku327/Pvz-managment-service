package models

import (
	"errors"
	"strconv"
	"time"
)

const (
	defaultPage = 1
	minPage     = 1

	defaultLimit = 10
	minLimit     = 1
	maxLimit     = 30

	datetimeLayout = "2006-01-02 15:04"
)

var (
	errInvalidStartDate = errors.New("invalid start date provided")
	errInvalidEndDate   = errors.New("invalid end date provided")
	errInvalidPage      = errors.New("invalid page provided")
	errInvalidLimit     = errors.New("invalid limit provided")
)

type FilterOption func(*FilterParams) error

type FilterParams struct {
	startDate time.Time
	endDate   time.Time
	page      int
	limit     int
}

func NewFilterParams(opts ...FilterOption) (*FilterParams, error) {
	params := &FilterParams{
		startDate: time.Time{},
		endDate:   time.Time{},
		page:      defaultPage,
		limit:     defaultLimit,
	}

	for _, opt := range opts {
		err := opt(params)
		if err != nil {
			return nil, err
		}
	}

	return params, nil
}

func WithStartDate(startDate string) FilterOption {
	return func(fp *FilterParams) error {
		if startDate == "" {
			return nil
		}

		d, err := time.Parse(datetimeLayout, startDate)
		if err != nil {
			return errInvalidStartDate
		}

		fp.startDate = d

		return nil
	}
}

func WithEndDate(endDate string) FilterOption {
	return func(fp *FilterParams) error {
		if endDate == "" {
			return nil
		}
		d, err := time.Parse(datetimeLayout, endDate)
		if err != nil {
			return errInvalidEndDate
		}

		fp.endDate = d

		return nil
	}
}

func WithPage(page string) FilterOption {
	return func(fp *FilterParams) error {
		if page == "" {
			return nil
		}

		p, err := strconv.Atoi(page)
		if err != nil || p < minPage {
			return errInvalidPage
		}

		fp.page = p

		return nil
	}
}

func WithLimit(limit string) FilterOption {
	return func(fp *FilterParams) error {
		if limit == "" {
			return nil
		}

		l, err := strconv.Atoi(limit)
		if err != nil || l > maxLimit || l < minLimit {
			return errInvalidLimit
		}

		fp.limit = l

		return nil
	}
}

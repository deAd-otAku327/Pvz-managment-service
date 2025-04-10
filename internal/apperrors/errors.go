package apperrors

import "errors"

var (
	ErrSmthWentWrong      = errors.New("something went wrong")
	ErrInvalidPvzID       = errors.New("invalid pvz id provided")
	ErrInvalidRole        = errors.New("invalid role provided")
	ErrInvalidCity        = errors.New("invalid city provided")
	ErrInvalidProductType = errors.New("invalid product type provided")
)

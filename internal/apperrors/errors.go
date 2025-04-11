package apperrors

import "errors"

var (
	ErrInvalidRequestBody   = errors.New("invalid request body provided")
	ErrInvalidRequestParams = errors.New("invalid request params provided")

	ErrSmthWentWrong      = errors.New("something went wrong")
	ErrInvalidPvzID       = errors.New("invalid pvz id provided")
	ErrInvalidRole        = errors.New("invalid role provided")
	ErrInvalidCity        = errors.New("invalid city provided")
	ErrInvalidProductType = errors.New("invalid product type provided")

	ErrInvalidStartDate = errors.New("invalid start date provided")
	ErrInvalidEndDate   = errors.New("invalid end date provided")
	ErrInvalidPage      = errors.New("invalid page provided")
	ErrInvalidLimit     = errors.New("invalid limit provided")
)

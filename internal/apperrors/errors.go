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

	ErrInvalidParamFormat = errors.New("invalid format of params provided")
	ErrInvalidPageParam   = errors.New("invalid page provided")
	ErrInvalidLimitParam  = errors.New("invalid limit provided")
)

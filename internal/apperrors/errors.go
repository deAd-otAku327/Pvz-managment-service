package apperrors

import "errors"

var (
	ErrInvalidRequestBody   = errors.New("invalid request body")
	ErrInvalidRequestParams = errors.New("invalid request params")

	ErrSmthWentWrong             = errors.New("something went wrong")
	ErrInvalidPvzID              = errors.New("invalid pvz id")
	ErrInvalidReceptionID        = errors.New("invalid reception id")
	ErrReceptionIsNotClosed      = errors.New("there is open reception in this pvz")
	ErrReceptionIsNotOpened      = errors.New("there is no opened reception in this pvz")
	ErrNoProductsInCurrReception = errors.New("there is no products in current reception in this pvz")
	ErrInvalidRole               = errors.New("invalid role")
	ErrInvalidCity               = errors.New("invalid city")
	ErrInvalidProductType        = errors.New("invalid product type")
	ErrInvalidReceptionStatus    = errors.New("invalid reception status")

	ErrInvalidParamFormat = errors.New("invalid format of params")
	ErrInvalidPageParam   = errors.New("invalid page")
	ErrInvalidLimitParam  = errors.New("invalid limit")
)

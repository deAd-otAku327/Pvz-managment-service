package werrors

// Package provide wrapping for errors.

type Werror interface {
	Error() string
	Code() int
}

type werror struct {
	Err        error
	StatusCode int
}

func New(err error, code int) Werror {
	return &werror{
		Err:        err,
		StatusCode: code,
	}
}

func (xe *werror) Error() string {
	return xe.Err.Error()
}

func (xe *werror) Code() int {
	return xe.StatusCode
}

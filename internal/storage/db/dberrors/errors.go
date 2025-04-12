package dberrors

import "errors"

var (
	ErrEnumTypeViolation   = errors.New("violated db enum type, inconsistensy detected")
	ErrForeignKeyViolation = errors.New("violated foreign key")
	ErrUpdateProhibited    = errors.New("update operation prohibited")
)

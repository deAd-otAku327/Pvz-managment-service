package dberrors

import "errors"

var (
	ErrEnumTypeViolation   = errors.New("violated db enum type, inconsistensy detected")
	ErrForeignKeyViolation = errors.New("violated foreign key")
	ErrInsertImpossible    = errors.New("insert operation impossible")
	ErrUpdateImpossible    = errors.New("update operation impossible")
)

package comerror

import "errors"

var (
	ErrDatabaseNotInit = errors.New("the database is not init")
	ErrUnknownError    = errors.New("unknown error")
)

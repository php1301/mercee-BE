package domain

import (
	"errors"
)

// Có thể coi như export các biến static

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen

	ErrInteralServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your Item is already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given Param is not valid")
)

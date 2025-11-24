package repository

import "errors"

// ErrNotFound is returned when a record is not found in the repository
var (
	ErrNotFound = errors.New("not found")
)

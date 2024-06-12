package domain

import "errors"

var (
	ErrInvalidCredentials  = errors.New("Invalid credentials")
	ErrInternalServerError = errors.New("Internal server error")
	ErrNotFound            = errors.New("Not found")
)

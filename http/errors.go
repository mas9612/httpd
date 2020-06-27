package http

import "errors"

var (
	// ErrInvalidRequest is the error corresponding status code 400.
	ErrInvalidRequest = errors.New("invalid request")
	// ErrNotFound is the error corresponding status code 404.
	ErrNotFound = errors.New("not found")
	// ErrInternalServerError is the error corresponding status code 500.
	ErrInternalServerError = errors.New("internal server error")
	// ErrMethodNotImplemented is the error corresponding status code 501.
	ErrMethodNotImplemented = errors.New("method not implemented")
)

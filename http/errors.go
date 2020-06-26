package http

import "errors"

var (
	ErrInvalidRequest       = errors.New("invalid request")
	ErrMethodNotImplemented = errors.New("method not implemented")
)

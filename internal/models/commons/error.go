package commons

import "errors"

var (
	ErrCredentials    = errors.New("invalid credentials")
	ErrNotfound       = errors.New("data not found")
	ErrConflict       = errors.New("data already exists")
	ErrInvalidToken   = errors.New("invalid token")
	ErrInternalServer = errors.New("internal server error")
	ErrInvalidInput   = errors.New("invalid input")
)

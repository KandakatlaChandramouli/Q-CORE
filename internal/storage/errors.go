package storage

import "errors"

var (
	ErrClosed          = errors.New("storage: closed")
	ErrBlockNotFound   = errors.New("storage: block not found")
	ErrAccountNotFound = errors.New("storage: account not found")
)

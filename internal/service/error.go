package service

import (
	"errors"
)

var (
	ErrInternal     = errors.New("internal error")
	ErrPostNotFound = errors.New("post not found")
)

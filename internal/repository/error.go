package repository

import "github.com/pkg/errors"

var (
	ErrInternal     = errors.New("internal error")
	ErrPostNotFound = errors.New("post not found")
)

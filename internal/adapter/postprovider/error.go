package postprovider

import "github.com/pkg/errors"

var (
	ErrRequestTimeout = errors.New("request timeout")
	ErrInternal       = errors.New("internal error")
)

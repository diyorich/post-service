package postfetcher

import (
	"errors"
	"fmt"
)

type FetchError struct {
	FetchErrs []error
	SaveErrs  []error
}

func (f *FetchError) Error() string {
	fErrs := errors.Join(f.FetchErrs...)
	sErrs := errors.Join(f.SaveErrs...)
	return fmt.Sprintf("fetchErrors: %v\n saveErrors: %v", fErrs.Error(), sErrs.Error())
}

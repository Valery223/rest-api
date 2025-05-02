package errdefs

import "errors"

var (
	// Common application errors
	ErrNotFound = errors.New("resource not found")
	// ErrInternalServer = errors.New("internal server error")
)

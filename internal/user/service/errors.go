package service

import "errors"

var (
	ErrEmailFormat     = errors.New("invalid email format")
	ErrPassowordFormat = errors.New("invalid password format")
)

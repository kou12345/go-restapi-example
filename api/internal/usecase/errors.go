package usecase

import "errors"

var (
	ErrInvalidUserID = errors.New("invalid user id")
	ErrInvalidAdmin  = errors.New("invalid admin")
)

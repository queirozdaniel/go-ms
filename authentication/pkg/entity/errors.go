package entity

import "errors"

var (
	ErrIDIsRequired          = errors.New("ID is required")
	ErrInvalidID             = errors.New("invalid ID")
	ErrEmailIsRequired       = errors.New("email is required")
	ErrNameIsRequired        = errors.New("name is required")
	ErrPasswordIsRequired    = errors.New("password is required")
	ErrExternalKeyIsRequired = errors.New("external key is required")
	ErrDescriptionIsRequired = errors.New("description is required")
	ErrKeyIsRequired         = errors.New("key is required")
)

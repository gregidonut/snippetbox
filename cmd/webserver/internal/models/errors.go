package models

import "errors"

var (
	ErrNoRecord           = errors.New("models: cannot find snippet")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

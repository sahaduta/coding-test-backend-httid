package apperror

import "errors"

var (
	// 400
	ErrInvalidInput = errors.New("invalid input")

	// 401
	ErrInvalidCred = errors.New("wrong username or password")
)

package apperror

import "errors"

var (
	// 400
	ErrInvalidInput = errors.New("invalid input")

	// 401
	ErrInvalidCred          = errors.New("wrong username or password")
	ErrInvalidCategoryId    = errors.New("invalid category id")
	ErrInvalidNewsArticleId = errors.New("invalid news article id")
	ErrInvalidCustomUrl     = errors.New("invalid custom url")

	ErrNotAuthorized   = errors.New("not authorized")
	ErrMissingMetadata = errors.New("missing metadata")

	// 404
	ErrCategoryIdNotFound    = errors.New("category id not found")
	ErrNewsArticleIdNotFound = errors.New("news article id not found")
	ErrCustomUrlNotFound     = errors.New("custom url not found")

	// 500
	ErrInternal = errors.New("internal error")
)

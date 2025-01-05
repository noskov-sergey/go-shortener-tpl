package shortener

import "errors"

var (
	ErrShortUrlNotFound = errors.New("can't found by short url")
)

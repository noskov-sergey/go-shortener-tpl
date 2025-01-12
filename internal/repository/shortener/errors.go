package shortener

import (
	"errors"

	"github.com/jackc/pgerrcode"
)

var (
	ErrShortURLNotFound = errors.New("can't found by short url")
	ErrNotUnique        = errors.New(pgerrcode.UniqueViolation)
)

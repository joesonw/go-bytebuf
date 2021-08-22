package bytebuf

import (
	"errors"
)

var (
	ErrBufferTypeNotMatch = errors.New("target buffer type is not desired")
	ErrInvalidLengthWrote = errors.New("wrote invalid length")
	ErrBufferReleased     = errors.New("buffer is already released")
)

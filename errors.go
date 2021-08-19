package go_bytebuf

import (
	"bytes"
	"errors"
	"io"
)

var (
	ErrNotEnoughBytes     = io.EOF
	ErrNotEnoughSpace     = bytes.ErrTooLarge
	ErrBufferTypeNotMatch = errors.New("target buffer type is not desired")
)

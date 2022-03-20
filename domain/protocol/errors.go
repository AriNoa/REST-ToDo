package protocol

import "errors"

var (
	ErrTagNotFound  = errors.New("tag not found")
	ErrTaskNotFound = errors.New("task not found")
	ErrInternal     = errors.New("internal error")
)

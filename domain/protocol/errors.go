package protocol

import "errors"

var (
	ErrTagNotFound       = errors.New("tag not found")
	ErrTaskNotFound      = errors.New("task not found")
	ErrNoTagName         = errors.New("no tag name error")
	ErrNoTaskTitle       = errors.New("no task title error")
	ErrTagAlreadyCreated = errors.New("tag already created")
	ErrInternal          = errors.New("internal error")
)

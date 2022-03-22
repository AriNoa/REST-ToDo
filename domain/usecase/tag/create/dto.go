package create

import (
	"github.com/AriNoa/REST-ToDo/domain/protocol"
)

type InputData struct {
	name string
}

func (i InputData) Name() string {
	return i.name
}

func NewInputData(name string) InputData {
	return InputData{
		name: name,
	}
}

type OutputData struct {
	tag protocol.Tag
}

func (o OutputData) Tag() protocol.Tag {
	return o.tag
}

func NewOutputData(tag protocol.Tag) OutputData {
	return OutputData{
		tag: tag,
	}
}

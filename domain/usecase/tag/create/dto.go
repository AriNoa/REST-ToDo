package create

import (
	"github.com/AriNoa/REST-ToDo/domain/protocol"
	"github.com/AriNoa/goutil/optional"
)

type InputData struct {
	name optional.Optional[string]
}

func (i InputData) Name() (string, bool) {
	return i.name.Get()
}

func NewInputData(
	name optional.Optional[string],
) InputData {
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

func NewOutputData(
	tag protocol.Tag,
) OutputData {
	return OutputData{
		tag: tag,
	}
}

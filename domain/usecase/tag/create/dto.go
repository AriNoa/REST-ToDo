package create

import (
	"github.com/AriNoa/REST-ToDo/domain/protocol"
	"github.com/AriNoa/goutil/optional"
)

type InputData struct {
	name optional.Optional[string]
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

func NewOutputData(
	tag protocol.Tag,
) OutputData {
	return OutputData{
		tag: tag,
	}
}

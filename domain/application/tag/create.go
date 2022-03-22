package tag

import (
	"github.com/AriNoa/REST-ToDo/domain/model/ids"
	model "github.com/AriNoa/REST-ToDo/domain/model/tag"
	"github.com/AriNoa/REST-ToDo/domain/protocol"
	usecase "github.com/AriNoa/REST-ToDo/domain/usecase/tag/create"
)

type CreateInteractor struct {
	repository model.Repository
}

func (m *CreateInteractor) Handle(data usecase.InputData) (usecase.OutputData, error) {
	name, err := model.NewName(data.Name())

	if err != nil {
		return usecase.OutputData{}, err
	}

	tag, err := m.repository.Create(name)

	if err != nil {
		return usecase.OutputData{}, err
	}

	return usecase.NewOutputData(
		protocol.NewTag(ids.String(tag.Id()), tag.Name().String()),
	), nil

}

func NewCreateInteractor(repository model.Repository) *CreateInteractor {
	return &CreateInteractor{
		repository: repository,
	}
}

package tag

import (
	model "github.com/AriNoa/REST-ToDo/domain/model/tag"
	usecase "github.com/AriNoa/REST-ToDo/domain/usecase/tag/create"
)

type CreateInteractor struct {
	repository model.Repository
}

func (m *CreateInteractor) Handle(data usecase.InputData) (usecase.OutputData, error) {
	return usecase.OutputData{}, nil
}

func NewCreateInteractor(repository model.Repository) *CreateInteractor {
	return &CreateInteractor{
		repository: repository,
	}
}

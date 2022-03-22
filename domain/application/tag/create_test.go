package tag

import (
	"testing"

	"github.com/AriNoa/REST-ToDo/domain/model/ids"
	model "github.com/AriNoa/REST-ToDo/domain/model/tag"
	"github.com/AriNoa/REST-ToDo/domain/protocol"
	"github.com/AriNoa/REST-ToDo/domain/usecase/tag/create"
	"github.com/AriNoa/goutil/functions"
	"github.com/stretchr/testify/assert"
)

func TestCreateTag(t *testing.T) {
	id := ids.New[ids.TagID]()
	name := functions.Must(model.NewName("hoge"))
	tag := functions.Must(model.New(id, name))

	repository := new(model.MockRepository)
	repository.On("Create", name).Return(tag, nil)

	interactor := NewCreateInteractor(repository)

	input := create.NewInputData("hoge")
	output, err := interactor.Handle(input)

	assert.NoError(t, err)
	assert.Equal(t, output.Tag(), protocol.NewTag(ids.String(id), name.String()))

	repository.AssertCalled(t, "Create", name)
}

func TestCreateNoNameTag(t *testing.T) {
	repository := new(model.MockRepository)
	interactor := NewCreateInteractor(repository)

	input := create.NewInputData("")
	_, err := interactor.Handle(input)

	assert.ErrorIs(t, err, protocol.ErrNoTagName)
}

package tag

import (
	"github.com/AriNoa/REST-ToDo/domain/model/ids"
	"github.com/stretchr/testify/mock"
)

type Repository interface {
	Create(name Name) (Tag, error)
	FindByID(id ids.TagID) (Tag, bool)
	FindByName(name Name) (Tag, bool)
	Update(tag Tag) error
	Delete(id ids.TagID) error
}

type MockRepository struct {
	mock.Mock
}

func (m MockRepository) Create(name Name) (Tag, error) {
	ret := m.Called(name)

	return ret.Get(0).(Tag), ret.Error(1)
}

func (m MockRepository) FindByID(id ids.TagID) (Tag, bool) {
	ret := m.Called(id)

	return ret.Get(0).(Tag), ret.Bool(1)
}

func (m MockRepository) FindByName(name Name) (Tag, bool) {
	ret := m.Called(name)

	return ret.Get(0).(Tag), ret.Bool(1)
}

func (m MockRepository) Update(tag Tag) error {
	ret := m.Called(tag)

	return ret.Error(0)
}

func (m MockRepository) Delete(id ids.TagID) error {
	ret := m.Called(id)

	return ret.Error(0)
}

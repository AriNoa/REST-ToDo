package create

import "github.com/stretchr/testify/mock"

type UseCase interface {
	Handle(InputData) (OutputData, error)
}

type MockUseCase struct {
	mock.Mock
}

func (m *MockUseCase) Handle(data InputData) (OutputData, error) {
	ret := m.Called(data)

	return ret.Get(0).(OutputData), ret.Error(1)
}

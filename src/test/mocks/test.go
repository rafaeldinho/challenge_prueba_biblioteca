// coverage:ignore
package mocks

import (
	"github.com/challenge_prueba_biblioteca/src/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

func (m *MockUseCase) GetCheck() model.Health {
	mocked := m.Called()
	return mocked.Get(0).(model.Health)
}

func MockHealthObject() model.Health {
	return model.Health{
		Status:  "UP",
		Version: "1.0.0",
	}
}

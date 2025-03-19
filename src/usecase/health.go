package usecase

import (
	"github.com/google/uuid"
	
	"github.com/challenge_prueba_biblioteca/src/domain/model"
)

type HealthUseCase interface {
	GetCheck() model.Health
}

type healthUseCase struct{}

func NewHealthUseCase() *healthUseCase {
	return &healthUseCase{}
}

func (h *healthUseCase) GetCheck() model.Health {
	return model.Health{
		Status:  "UP",
		Version: uuid.New().String(),
	}
}

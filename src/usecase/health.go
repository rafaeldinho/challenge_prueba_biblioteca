package usecase

import (
	"github.com/challenge_prueba_biblioteca/src/domain/model"
	"github.com/challenge_prueba_biblioteca/src/shared"
	log "github.com/sirupsen/logrus"
	"github.com/google/uuid"
)

type HealthUseCase interface {
	GetCheck() model.Health
}

type healthUseCase struct{}

var logger = log.WithFields(log.Fields{
	"layer": shared.HealthLayer,
})

func NewHealthUseCase() *healthUseCase {
	return &healthUseCase{}
}

func (h *healthUseCase) GetCheck() model.Health {
	logger.Info("getting service status")
	return model.Health{
		Status:  "UP",
		Version:  uuid.New().String(),
	}
}

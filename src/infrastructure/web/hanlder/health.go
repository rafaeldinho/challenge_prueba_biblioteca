package handler

import (
	"net/http"

	"github.com/challenge_prueba_biblioteca/src/usecase"
	"github.com/labstack/echo"
)

type healthHandler struct {
	useCae usecase.HealthUseCase
}

func NewHealthHandler(e *echo.Echo, useCase usecase.HealthUseCase) *healthHandler {
	h := &healthHandler{useCase}
	e.GET("/health", h.HealthCheck)
	return h
}

func (h *healthHandler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, h.useCae.GetCheck())
}

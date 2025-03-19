package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/challenge_prueba_biblioteca/src/test/mocks"
)

const jsonHealthCheck = "{\"status\":\"UP\",\"version\":\"1.0.0\"}\n"

func TestNewHealthHandler(t *testing.T) {

	t.Run("When response health check is up", func(t *testing.T) {

		useCase := new(mocks.MockHealthUseCase)
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)

		useCase.On("GetCheck").Return(mocks.MockHealthObject())

		healthCheck := NewHealthHandler(e, useCase)

		_ = healthCheck.HealthCheck(echoContext)

		assert.Equal(t, jsonHealthCheck, rec.Body.String())
		assert.Equal(t, http.StatusOK, rec.Code)
		useCase.AssertCalled(t, "GetCheck")
	})
}

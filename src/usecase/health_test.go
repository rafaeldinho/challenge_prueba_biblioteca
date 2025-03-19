package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHealthUseCase(t *testing.T) {

	t.Run("When calling GetCheck should return health status object", func(t *testing.T) {

		resutl := NewHealthUseCase().GetCheck()

		assert.NotEmpty(t, resutl.Status)
		assert.NotEmpty(t, resutl.Version)
		assert.NotNil(t, resutl)
		assert.IsType(t, "", resutl.Status)
		assert.IsType(t, "", resutl.Version)
	})
}

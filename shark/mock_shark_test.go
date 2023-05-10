package shark

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"integrationtests/prey"
)

func TestMockSharkHunt(t *testing.T){
	// Arrange
	shark := NewMockShark()
	shark.HuntFunc = func(prey prey.Prey) error {
		return nil
	}
	tuna := prey.NewStubPrey()
	tuna.GetSpeedFunc = func() float64 {
		return 5.0
	}

	// Act
	err := shark.Hunt(tuna)

	// Assert
	assert.NoError(t, err)
}


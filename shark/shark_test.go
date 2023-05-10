package shark

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"integrationtests/prey"
	"integrationtests/simulator"
	"integrationtests/pkg/storage"
)

func TestSharkHunt(t *testing.T){
	t.Run("the shark succedeed in hunting its prey", func(t *testing.T) {
		// Arrange
		sharkStorage := storage.NewMockStorageTestify()
		sharkStorage.Mock.On("GetValue", "white_shark_speed").Return(100.0)
		sharkStorage.Mock.On("GetValue", "white_shark_x").Return(2.0)
		sharkStorage.Mock.On("GetValue", "white_shark_y").Return(2.0)
		shark := CreateWhiteShark(simulator.NewCatchSimulator(500.0), sharkStorage)
		tunaStorage := storage.NewMockStorageTestify()
		tunaStorage.Mock.On("GetValue", "tuna_speed").Return(5.0)
		tuna := prey.CreateTuna(tunaStorage)
		// Act
		err := shark.Hunt(tuna)

		// Assert
		assert.NoError(t, err)
	})
	t.Run("the shark couldnt catch its prey", func(t *testing.T) {
		// Arrange
		sharkStorage := storage.NewMockStorageTestify()
		sharkStorage.Mock.On("GetValue", "white_shark_speed").Return(30.0)
		sharkStorage.Mock.On("GetValue", "white_shark_x").Return(0.0)
		sharkStorage.Mock.On("GetValue", "white_shark_y").Return(0.0)
		shark := CreateWhiteShark(simulator.NewCatchSimulator(30.0), sharkStorage)
		tunaStorage := storage.NewMockStorageTestify()
		tunaStorage.Mock.On("GetValue", "tuna_speed").Return(50.0)
		tuna := prey.CreateTuna(tunaStorage)
		// Act
		err := shark.Hunt(tuna)

		// Assert
		assert.Error(t, err)
	})
}
package storage

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetValue(t *testing.T) {
	t.Run("returns value from db", func(t *testing.T) {
		// Arrange
		storageMock := NewMockStorageTestify()
		storageMock.Mock.On("GetValue", "speed").Return(5.0)

		// Act
		result := storageMock.GetValue("speed")

		// Assert
		assert.Equal(t, 5.0, result)
	})
	t.Run("returns nil if key does not exist in db", func(t *testing.T) {
		// Arrange
		storageMock := NewMockStorageTestify()
		storageMock.Mock.On("GetValue", "speed").Return(nil)

		// Act
		result := storageMock.GetValue("speed")

		// Assert
		assert.Nil(t, result)
	})
}
package shark

import (
	"integrationtests/prey"
)

type MockShark struct {
	HuntFunc func(prey.Prey) error
	HuntSpy  bool
}

func NewMockShark() *MockShark {
	return &MockShark{}
}

func (mock *MockShark) Hunt(prey prey.Prey) error {
	mock.HuntSpy = true
	return mock.HuntFunc(prey)
}
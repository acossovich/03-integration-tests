package storage

import (
	"github.com/stretchr/testify/mock"
)

type MockStorageTestify struct {
	Mock *mock.Mock
}

func NewMockStorageTestify() *MockStorageTestify {
	return &MockStorageTestify{Mock: &mock.Mock{}}
}

func (mock *MockStorageTestify) GetValue(key string) interface{} {
	args := mock.Mock.Called(key)
	return args.Get(0)
}
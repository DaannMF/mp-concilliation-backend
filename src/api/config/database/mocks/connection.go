package mocks

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockGormConnection struct {
	mock.Mock
}

func (mock *MockGormConnection) Connect() (client *gorm.DB, err error) {
	responseArgs := mock.Called()
	response := responseArgs.Get(0)
	err = responseArgs.Error(1)
	if response != nil {
		return response.(*gorm.DB), err
	}

	return nil, err
}

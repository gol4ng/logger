// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import logger "github.com/gol4ng/logger"
import mock "github.com/stretchr/testify/mock"

// ErrorHandler is an autogenerated mock type for the ErrorHandler type
type ErrorHandler struct {
	mock.Mock
}

// HandleError provides a mock function with given fields: error, entry
func (_m *ErrorHandler) HandleError(error error, entry logger.Entry) {
	_m.Called(error, entry)
}
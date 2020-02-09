package testing_test

import (
	"errors"
	"testing"

	"github.com/gol4ng/logger"
	testing2 "github.com/gol4ng/logger/testing"
	"github.com/stretchr/testify/assert"
)

type testingTMock struct {
	callback func(format string, args ...interface{})
}

func (t testingTMock) Errorf(format string, args ...interface{}) {
	t.callback(format, args...)
}

func Test_AssertErrorHandlerNotCalled(t *testing.T) {
	testingTMock := testingTMock{
		callback: func(format string, args ...interface{}) {
			assert.Equal(t, "\n%s", format)
			assert.Contains(t, args[0], "ErrorHandler must not be called")
			assert.Contains(t, args[0], `unexpected error "my_fake_error_message" occurred for logger.Entry{Message:"", Level:0, Context:<nil>}`)
		},
	}
	assertionFunction := testing2.AssertErrorHandlerNotCalled(testingTMock)

	assertionFunction(errors.New("my_fake_error_message"), logger.Entry{})
}

func Test_AssertErrorHandlerNotCalled_ShouldNotBeCalled(t *testing.T) {
	testingTMock := testingTMock{
		callback: func(format string, args ...interface{}) {
			assert.Fail(t, "testingMock should not be called")
		},
	}
	testing2.AssertErrorHandlerNotCalled(testingTMock)
}

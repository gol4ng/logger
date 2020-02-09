package testing

import (
	"github.com/gol4ng/logger"
	"github.com/stretchr/testify/assert"
)

func AssertErrorHandlerNotCalled(t assert.TestingT) func(err error, entry logger.Entry) {
	return func(err error, entry logger.Entry) {
		assert.Fail(t, "ErrorHandler must not be called", "unexpected error \"%s\" occurred for %#v", err, entry)
	}
}

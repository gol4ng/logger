package testing

import (
	"testing"

	"github.com/gol4ng/logger"
	"github.com/stretchr/testify/assert"
)

func AssertErrorHandlerNotCalled(t *testing.T) func(err error, entry logger.Entry) {
	return func(err error, entry logger.Entry) {
		assert.Fail(t, "ErrorHandler must not be called", "unexpected error \"%s\" occurred for %#v", err, entry)
	}
}

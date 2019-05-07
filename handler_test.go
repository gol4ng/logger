package logger_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
)

func TestNopHandler_Handle(t *testing.T) {
	assert.Nil(t, logger.NopHandler(logger.Entry{}))
}

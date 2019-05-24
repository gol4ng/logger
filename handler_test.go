package logger_test

import (
	"testing"

	"github.com/gol4ng/logger"
	"github.com/stretchr/testify/assert"
)

func TestNopHandler_Handle(t *testing.T) {
	assert.Nil(t, logger.NopHandler(logger.Entry{}))
}

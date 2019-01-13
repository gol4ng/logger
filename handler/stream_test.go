package handler_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/handler"
)

func TestNilStream_Handle(t *testing.T) {
	assert.Nil(t, handler.NewNilStream().Handle(logger.Entry{}))
}

func TestStream_Handle(t *testing.T) {
	var b bytes.Buffer
	h := handler.NewStream(&b, logger.NewNilFormatter())

	assert.Nil(t, h.Handle(logger.Entry{Message: "test message", Level: logger.WarnLevel, Context: nil}))
	assert.Equal(t, "{test message warn <nil>}\n", b.String())
}

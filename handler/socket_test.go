package handler_test

import (
	"net"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSocket(t *testing.T) {
	Connection := &net.TCPConn{}
	patch := gomonkey.NewPatches()
	patch.ApplyMethod(reflect.TypeOf(Connection), "Write", func(conn *net.TCPConn, b []byte) (n int, err error) {
		assert.Equal(t, []uint8("my formatter return"), b)
		return 99, nil
	})
	defer patch.Reset()

	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.On("Format", mock.AnythingOfType("logger.Entry")).Return("my formatter return")

	h := handler.Socket(Connection, &mockFormatter)

	assert.Nil(t, h(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
}

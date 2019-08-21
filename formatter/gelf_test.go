package formatter_test

import (
	"errors"
	"os"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
)

func TestNewGelf_WillPanic(t *testing.T) {
	err := errors.New("my_hostname_error")
	monkey.Patch(os.Hostname, func() (string, error) { return "", err })
	assert.PanicsWithValue(t, err, func() { formatter.NewGelf() })
}

func TestGelf_Format(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	monkey.Patch(os.Hostname, func() (string, error) { return "my_fake_hostname", nil })
	defer monkey.UnpatchAll()

	gelf := formatter.NewGelf()

	assert.Equal(
		t,
		"{\"version\":\"1.1\",\"host\":\"my_fake_hostname\",\"level\":6,\"timestamp\":513216000.000,\"short_message\":\"My log message\",\"full_message\":\"<info> My log message [ <my key:my_value> ]\",\"_my_key\":\"my_value\"}",
		gelf.Format(logger.Entry{Message: "My log message", Level: logger.InfoLevel, Context: logger.NewContext().Add("my key", "my_value")}),
	)
}

func TestGelfTCP_Format(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	monkey.Patch(os.Hostname, func() (string, error) { return "my_fake_hostname", nil })
	defer monkey.UnpatchAll()

	gelf := formatter.NewGelfTCP()

	assert.Equal(
		t,
		"{\"version\":\"1.1\",\"host\":\"my_fake_hostname\",\"level\":6,\"timestamp\":513216000.000,\"short_message\":\"My log message\",\"full_message\":\"<info> My log message [ <my key:my_value> ]\",\"_my_key\":\"my_value\"}\x00",
		gelf.Format(logger.Entry{Message: "My log message", Level: logger.InfoLevel, Context: logger.NewContext().Add("my key", "my_value")}),
	)
}

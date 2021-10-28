package formatter_test

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/stretchr/testify/assert"
)

func TestNewGelf_WillPanic(t *testing.T) {
	err := errors.New("my_hostname_error")
	patch := gomonkey.NewPatches()
	patch.ApplyFunc(os.Hostname, func() (string, error) { return "", err })
	defer patch.Reset()
	assert.PanicsWithValue(t, err, func() { formatter.NewGelf() })
}

func TestGelf_Format(t *testing.T) {
	patch := gomonkey.NewPatches()
	patch.ApplyFunc(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	patch.ApplyFunc(os.Hostname, func() (string, error) { return "my_fake_hostname", nil })
	defer patch.Reset()

	gelf := formatter.NewGelf()

	assert.Equal(
		t,
		"{\"version\":\"1.1\",\"host\":\"my_fake_hostname\",\"level\":6,\"timestamp\":513216000.000,\"short_message\":\"My log message\",\"full_message\":\"<info> My log message [ <my key:my_value> ]\",\"_my_key\":\"my_value\"}",
		gelf.Format(logger.Entry{Message: "My log message", Level: logger.InfoLevel, Context: logger.NewContext().Add("my key", "my_value")}),
	)
}

func TestGelfTCP_Format(t *testing.T) {
	patch := gomonkey.NewPatches()
	patch.ApplyFunc(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	patch.ApplyFunc(os.Hostname, func() (string, error) { return "my_fake_hostname", nil })
	defer patch.Reset()

	gelf := formatter.NewGelfTCP()

	assert.Equal(
		t,
		"{\"version\":\"1.1\",\"host\":\"my_fake_hostname\",\"level\":6,\"timestamp\":513216000.000,\"short_message\":\"My log message\",\"full_message\":\"<info> My log message [ <my key:my_value> ]\",\"_my_key\":\"my_value\"}\x00",
		gelf.Format(logger.Entry{Message: "My log message", Level: logger.InfoLevel, Context: logger.NewContext().Add("my key", "my_value")}),
	)
}

// =====================================================================================================================
// ==================================================== EXAMPLES =======================================================
// =====================================================================================================================

func ExampleGelf_Format() {
	patch := gomonkey.NewPatches()
	patch.ApplyFunc(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	patch.ApplyFunc(os.Hostname, func() (string, error) { return "my_fake_hostname", nil })
	defer patch.Reset()

	gelfFormatter := formatter.NewGelf()

	fmt.Println(gelfFormatter.Format(logger.Entry{Message: "My log message", Level: logger.InfoLevel, Context: logger.NewContext().Add("my key", "my_value")}))

	//Output:
	//{"version":"1.1","host":"my_fake_hostname","level":6,"timestamp":513216000.000,"short_message":"My log message","full_message":"<info> My log message [ <my key:my_value> ]","_my_key":"my_value"}
}

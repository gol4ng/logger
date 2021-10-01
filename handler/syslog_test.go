// +build !windows,!plan9,!js

package handler_test

import (
	"errors"
	"log/syslog"
	"reflect"
	"testing"

	"bou.ke/monkey"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/mocks"
)

func TestSyslog_HandleWithDialError(t *testing.T) {
	monkey.Patch(syslog.Dial, func(network, raddr string, priority syslog.Priority, tag string) (*syslog.Writer, error) {
		assert.Equal(t, "fake_network", network)
		assert.Equal(t, "fake_raddr", raddr)
		assert.Equal(t, syslog.LOG_DEBUG, priority)
		assert.Equal(t, "fake_tag", tag)
		return nil, errors.New("fake_syslog_dial_error")
	})
	defer monkey.UnpatchAll()

	logEntry := logger.Entry{Level: -1}

	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.AssertNotCalled(t, "Format", logEntry)

	h, err := handler.Syslog(&mockFormatter, "fake_network", "fake_raddr", syslog.LOG_DEBUG, "fake_tag")

	assert.Error(t, err, "fake_syslog_dial_error")
	assert.Nil(t, h)
}

func TestSyslog_HandleWithWriteError(t *testing.T) {
	var w *syslog.Writer // Has to be a pointer to because `Dial` has a pointer receiver

	monkey.Patch(syslog.Dial, func(network, raddr string, priority syslog.Priority, tag string) (*syslog.Writer, error) {
		assert.Equal(t, "fake_network", network)
		assert.Equal(t, "fake_raddr", raddr)
		assert.Equal(t, syslog.LOG_DEBUG, priority)
		assert.Equal(t, "fake_tag", tag)
		return w, nil
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(w), "Write", func(_ *syslog.Writer, m []byte) (int, error) {
		assert.Equal(t, []byte("fake_syslog_message"), m)
		return 0, errors.New("fake_syslog_write_error")
	})
	defer monkey.UnpatchAll()

	logEntry := logger.Entry{Level: -1}

	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.On("Format", logEntry).Return("fake_syslog_message")

	h, _ := handler.Syslog(&mockFormatter, "fake_network", "fake_raddr", syslog.LOG_DEBUG, "fake_tag")

	assert.EqualError(t, h(logEntry), "fake_syslog_write_error")
}

func TestSyslog_Handle(t *testing.T) {
	syslogMethods := map[logger.Level]string{
		logger.DebugLevel:     "Debug",
		logger.NoticeLevel:    "Notice",
		logger.InfoLevel:      "Info",
		logger.WarningLevel:   "Warning",
		logger.ErrorLevel:     "Err",
		logger.AlertLevel:     "Alert",
		logger.CriticalLevel:  "Crit",
		logger.EmergencyLevel: "Emerg",
	}

	tests := []struct {
		name             string
		level            logger.Level
		methodToBeCalled string
	}{
		{level: logger.DebugLevel},
		{level: logger.InfoLevel},
		{level: logger.NoticeLevel},
		{level: logger.WarningLevel},
		{level: logger.ErrorLevel},
		{level: logger.AlertLevel},
		{level: logger.CriticalLevel},
		{level: logger.EmergencyLevel},
	}
	for _, test := range tests {
		t.Run("test syslog "+test.level.String(), func(t *testing.T) {
			syslogMsg := "this is a test " + test.level.String() + " syslog message"
			var w *syslog.Writer // Has to be a pointer to because `Dial` has a pointer receiver

			monkey.Patch(syslog.Dial, func(network, raddr string, priority syslog.Priority, tag string) (*syslog.Writer, error) {
				assert.Equal(t, "fake_network", network)
				assert.Equal(t, "fake_raddr", raddr)
				assert.Equal(t, syslog.LOG_DEBUG, priority)
				assert.Equal(t, "golang", tag)
				return w, nil
			})

			for l, syslogMethod := range syslogMethods {
				if test.level != l {
					monkey.PatchInstanceMethod(reflect.TypeOf(w), syslogMethod, func(_ *syslog.Writer, m string) error {
						t.Fatal("method syslog.Writer::" + syslogMethod + " was not expected to be called")
						return nil
					})
					continue
				}
				monkey.PatchInstanceMethod(reflect.TypeOf(w), syslogMethod, func(_ *syslog.Writer, m string) error {
					assert.Equal(t, syslogMsg, m)
					return nil
				})
			}
			defer monkey.UnpatchAll()

			//LOG DEBUG HERE
			logEntry := logger.Entry{Level: test.level}

			mockFormatter := mocks.FormatterInterface{}
			mockFormatter.On("Format", logEntry).Return(syslogMsg)

			h, _ := handler.Syslog(&mockFormatter, "fake_network", "fake_raddr", syslog.LOG_DEBUG, "")

			assert.Nil(t, h(logEntry))
		})
	}
}

// =====================================================================================================================
// ================================================= EXAMPLES ==========================================================
// =====================================================================================================================

// You can run the command below to show syslog messages
// syslog -F '$Time $Host $(Sender)[$(PID)] <$((Level)(str))>: $Message'
//Jan 22 22:42:14 hades my_go_logger[113] <Notice>: notice Log example2 &map[ctx_key:ctx_value]
//Jan 22 22:42:14 hades my_go_logger[113] <Warning>: warning Log example3 &map[ctx_key:ctx_value]
//Jan 22 22:42:14 hades my_go_logger[113] <Error>: error Log example4 &map[ctx_key:ctx_value]
//Jan 22 22:42:14 hades my_go_logger[113] <Critical>: critical Log example5 &map[ctx_key:ctx_value]
//Jan 22 22:42:14 hades my_go_logger[113] <Alert>: alert Log example6 &map[ctx_key:ctx_value]
//Jan 22 22:42:14 hades my_go_logger[113] <Emergency>: emergency Log example7 &map[ctx_key:ctx_value]
func ExampleSyslog() {
	syslogHandler, _ := handler.Syslog(
		formatter.NewLine("%[2]s %[1]s %[3]s"),
		"",
		"",
		syslog.LOG_DEBUG,
		"my_go_logger")

	_ = syslogHandler(logger.Entry{Message: "Log example", Level: logger.DebugLevel})
	_ = syslogHandler(logger.Entry{Message: "Log example", Level: logger.InfoLevel})
	_ = syslogHandler(logger.Entry{Message: "Log example", Level: logger.NoticeLevel})
	_ = syslogHandler(logger.Entry{Message: "Log example", Level: logger.WarningLevel})
	_ = syslogHandler(logger.Entry{Message: "Log example", Level: logger.ErrorLevel})
	_ = syslogHandler(logger.Entry{Message: "Log example", Level: logger.CriticalLevel})
	_ = syslogHandler(logger.Entry{Message: "Log example", Level: logger.AlertLevel})
	_ = syslogHandler(logger.Entry{Message: "Log example", Level: logger.EmergencyLevel})
	//Output:
}

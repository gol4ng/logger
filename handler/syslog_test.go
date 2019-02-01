package handler_test

import (
	"bou.ke/monkey"
	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"log/syslog"
	"reflect"
	"testing"
)

func TestSyslog_HandleWithError(t *testing.T) {
	var w *syslog.Writer // Has to be a pointer to because `Dial` has a pointer receiver

	monkey.Patch(syslog.Dial, func(network, raddr string, priority syslog.Priority, tag string) (*syslog.Writer, error) {
		assert.Equal(t,"fake_network", network)
		assert.Equal(t,"fake_raddr", raddr)
		assert.Equal(t, syslog.LOG_DEBUG, priority)
		assert.Equal(t,"fake_tag", tag)
		return w, nil
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(w), "Write", func(_ *syslog.Writer, m []byte) (int, error) {
		assert.Equal(t, []byte("this is a test syslog message"), m)
		return 0, errors.New("error occurred while calling syslog.Writer::Write")
	})
	defer monkey.UnpatchAll()

	logEntry := logger.Entry{Level: 100}

	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.On("Format", logEntry).Return("this is a test syslog message")

	h, _ := handler.NewSyslog(&mockFormatter, "fake_network", "fake_raddr", syslog.LOG_DEBUG, "fake_tag")

	assert.EqualError(t, h.Handle(logEntry), "error occurred while calling syslog.Writer::Write")
}

func TestSyslog_Handle(t *testing.T) {
	syslogMethods := map[logger.Level]string{
		logger.DebugLevel: "Debug",
		logger.NoticeLevel: "Notice",
		logger.InfoLevel: "Info",
		logger.WarningLevel: "Warning",
		logger.ErrorLevel: "Err",
		logger.AlertLevel: "Alert",
		logger.CriticalLevel: "Crit",
		logger.EmergencyLevel: "Emerg",
	}

	tests := []struct {
		name string
		level logger.Level
		methodToBeCalled string
	}{
		{ name: "test syslog debug", level: logger.DebugLevel, methodToBeCalled: "Debug" },
		{ name: "test syslog info", level: logger.InfoLevel, methodToBeCalled: "Info" },
		{ name: "test syslog notice", level: logger.NoticeLevel, methodToBeCalled: "Notice" },
		{ name: "test syslog warning", level: logger.WarningLevel, methodToBeCalled: "Warning" },
		{ name: "test syslog error", level: logger.ErrorLevel, methodToBeCalled: "Err" },
		{ name: "test syslog alert", level: logger.AlertLevel, methodToBeCalled: "Alert" },
		{ name: "test syslog critical", level: logger.CriticalLevel, methodToBeCalled: "Crit" },
		{ name: "test syslog emergency", level: logger.EmergencyLevel, methodToBeCalled: "Emerg" },
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var w *syslog.Writer // Has to be a pointer to because `Dial` has a pointer receiver

			monkey.Patch(syslog.Dial, func(network, raddr string, priority syslog.Priority, tag string) (*syslog.Writer, error) {
				assert.Equal(t,"fake_network", network)
				assert.Equal(t,"fake_raddr", raddr)
				assert.Equal(t, syslog.LOG_DEBUG, priority)
				assert.Equal(t,"golang", tag)
				return w, nil
			})

			for _, syslogMethod := range syslogMethods {
				if test.methodToBeCalled != syslogMethod {
					monkey.PatchInstanceMethod(reflect.TypeOf(w), syslogMethod, func(_ *syslog.Writer, m string) error {
						t.Fatal("method syslog.Writer::"+syslogMethod+" was not expected to be called")
						return nil
					})
					continue
				}
				monkey.PatchInstanceMethod(reflect.TypeOf(w), syslogMethod, func(_ *syslog.Writer, m string) error {
					assert.Equal(t, "this is a test "+ test.level.String() +" syslog message", m)
					return nil
				})
			}
			defer monkey.UnpatchAll()

			//LOG DEBUG HERE
			logEntry := logger.Entry{Level: test.level}

			mockFormatter := mocks.FormatterInterface{}
			mockFormatter.On("Format", logEntry).Return("this is a test "+test.level.String()+" syslog message")

			h, _ := handler.NewSyslog(&mockFormatter, "fake_network", "fake_raddr", syslog.LOG_DEBUG, "")

			assert.Nil(t, h.Handle(logEntry))
		})
	}
}

package logger_test

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"log/syslog"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gol4ng/logger/mocks"
	testing_logger "github.com/gol4ng/logger/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/middleware"
)

func TestLogger_ErrorHandler(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{})
	log.SetOutput(buffer)

	err := errors.New("my error")
	l := logger.NewLogger(func(entry logger.Entry) error {
		return err
	})
	l.Debug("my_fake_message")

	assert.Regexp(t, ".* my error {my_fake_message debug <nil>}\n", buffer.String())
}

func TestLogger_Log(t *testing.T) {
	tests := []struct {
		name  string
		level logger.Level
	}{
		{
			name:  "test Debug(...)",
			level: logger.DebugLevel,
		},
		{
			name:  "test Info(...)",
			level: logger.InfoLevel,
		},
		{
			name:  "test Notice(...)",
			level: logger.NoticeLevel,
		},
		{
			name:  "test Warning(...)",
			level: logger.WarningLevel,
		},
		{
			name:  "test Error(...)",
			level: logger.ErrorLevel,
		},
		{
			name:  "test Critical(...)",
			level: logger.CriticalLevel,
		},
		{
			name:  "test Alert(...)",
			level: logger.AlertLevel,
		},
		{
			name:  "test Emergency(...)",
			level: logger.EmergencyLevel,
		},
		{
			name:  "test Log(custom levelString)",
			level: logger.Level(127),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logger.NewLogger(func(entry logger.Entry) error {
				assert.Equal(t, "l message", entry.Message)
				assert.Equal(t, tt.level, entry.Level)
				assert.Equal(t, logger.Context{}, *entry.Context)

				return nil
			})
			l.ErrorHandler = testing_logger.AssertErrorHandlerNotCalled(t)

			switch tt.level {
			case logger.DebugLevel:
				l.Debug("l message")
			case logger.InfoLevel:
				l.Info("l message")
			case logger.NoticeLevel:
				l.Notice("l message")
			case logger.WarningLevel:
				l.Warning("l message")
			case logger.ErrorLevel:
				l.Error("l message")
			case logger.CriticalLevel:
				l.Critical("l message")
			case logger.AlertLevel:
				l.Alert("l message")
			case logger.EmergencyLevel:
				l.Emergency("l message")
			default:
				l.Log("l message", tt.level)
			}
		})
	}
}

func TestNewLogger_WillPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			assert.Equal(t, err, errors.New("handler must not be <nil>"))
		}
	}()
	logger.NewLogger(nil)
}

func TestNewLogger_LogWithError(t *testing.T) {
	err := errors.New("my error")
	l := logger.NewLogger(func(entry logger.Entry) error {
		return err
	})

	errorHandler := mocks.ErrorHandler{}
	l.ErrorHandler = errorHandler.HandleError
	errorHandler.On("HandleError", err, mock.AnythingOfType("logger.Entry"))

	l.Debug("l message")
	l.Info("l message")
	l.Notice("l message")
	l.Warning("l message")
	l.Error("l message")
	l.Critical("l message")
	l.Alert("l message")
	l.Emergency("l message")
	l.Log("l message", logger.Level(127))

	errorHandler.AssertCalled(t, "HandleError", err, mock.AnythingOfType("logger.Entry"))
	errorHandler.AssertExpectations(t)
}

func TestNewNilLogger_Log(t *testing.T) {
	l := logger.NewNopLogger()

	l.Debug("l message")
	l.Info("l message")
	l.Notice("l message")
	l.Warning("l message")
	l.Error("l message")
	l.Critical("l message")
	l.Alert("l message")
	l.Emergency("l message")
	l.Log("l message", logger.Level(127))

	l.ErrorHandler = testing_logger.AssertErrorHandlerNotCalled(t)
}

func TestNewLogger_Wrap(t *testing.T) {
	i := 0

	mockHandlerCalled := false
	mockHandler := func(entry logger.Entry) error {
		mockHandlerCalled = true
		assert.Equal(t, "l message", entry.Message)
		assert.Equal(t, logger.DebugLevel, entry.Level)
		assert.Equal(t, logger.Context{}, *entry.Context)
		assert.Equal(t, 3, i)

		i++
		return nil
	}

	l := logger.NewLogger(mockHandler)
	l.ErrorHandler = testing_logger.AssertErrorHandlerNotCalled(t)

	countingMiddleware := func(expectedI int) logger.MiddlewareInterface {
		return func(h logger.HandlerInterface) logger.HandlerInterface {
			return func(entry logger.Entry) error {
				assert.Equal(t, expectedI, i)
				i++
				return h(entry)
			}
		}
	}

	l.Wrap(
		countingMiddleware(2),
		countingMiddleware(1),
		countingMiddleware(0),
	)

	l.Debug("l message")

	assert.True(t, mockHandlerCalled)
	assert.Equal(t, 4, i)
}

func TestNewLogger_WrapNew(t *testing.T) {
	i := 0

	mockHandlerCalled := false
	mockHandlerExpectedI := 0
	mockHandler := func(entry logger.Entry) error {
		mockHandlerCalled = true
		assert.Equal(t, "l message", entry.Message)
		assert.Equal(t, logger.DebugLevel, entry.Level)
		assert.Equal(t, logger.Context{}, *entry.Context)
		assert.Equal(t, mockHandlerExpectedI, i)

		i++
		return nil
	}

	l := logger.NewLogger(mockHandler)
	l.ErrorHandler = testing_logger.AssertErrorHandlerNotCalled(t)

	countingMiddleware := func(expectedI int) logger.MiddlewareInterface {
		return func(h logger.HandlerInterface) logger.HandlerInterface {
			return func(entry logger.Entry) error {
				assert.Equal(t, expectedI, i)
				i++
				return h(entry)
			}
		}
	}

	l2 := l.WrapNew(
		countingMiddleware(2),
		countingMiddleware(1),
		countingMiddleware(0),
	)

	mockHandlerExpectedI = 0
	l.Debug("l message")
	assert.True(t, mockHandlerCalled)
	assert.Equal(t, 1, i)

	i = 0
	mockHandlerCalled = false
	mockHandlerExpectedI = 3

	l2.Debug("l message")
	assert.True(t, mockHandlerCalled)
	assert.Equal(t, 4, i)
}

// =====================================================================================================================
// ================================================= EXAMPLES ==========================================================
// =====================================================================================================================

func ExampleLogger_callerHandler() {
	output := &Output{}
	myLogger := logger.NewLogger(
		middleware.Caller(3)(
			handler.Stream(output, formatter.NewLine("lvl: %[2]s | msg: %[1]s | ctx: %[3]v")),
		),
	)

	myLogger.Debug("Log example")
	myLogger.Info("Log example")
	myLogger.Notice("Log example")
	myLogger.Warning("Log example")
	myLogger.Error("Log example")
	myLogger.Critical("Log example")
	myLogger.Alert("Log example")
	myLogger.Emergency("Log example")

	output.Contains([]string{
		"lvl: debug | msg: Log example | ctx:", "<file:/", "<line:",
		"lvl: info | msg: Log example | ctx:", "<file:/", "<line:",
		"lvl: notice | msg: Log example | ctx:", "<file:/", "<line:",
		"lvl: warning | msg: Log example | ctx:", "<file:/", "<line:",
		"lvl: error | msg: Log example | ctx:", "<file:/", "<line:",
		"lvl: critical | msg: Log example | ctx:", "<file:/", "<line:",
		"lvl: alert | msg: Log example | ctx:", "<file:/", "<line:",
		"lvl: emergency | msg: Log example | ctx:", "<file:/", "<line:",
	})

	//Output:
}

func ExampleLogger_lineFormatter() {
	output := &Output{}
	myLogger := logger.NewLogger(
		handler.Stream(output, formatter.NewLine("lvl: %[2]s | msg: %[1]s | ctx: %[3]v")),
	)

	myLogger.Debug("Log example", logger.Any("my_key", "my_value"))
	myLogger.Info("Log example", logger.Any("my_key", "my_value"))
	myLogger.Notice("Log example", logger.Any("my_key", "my_value"))
	myLogger.Warning("Log example", logger.Any("my_key", "my_value"))
	myLogger.Error("Log example", logger.Any("my_key", "my_value"))
	myLogger.Critical("Log example", logger.Any("my_key", "my_value"))
	myLogger.Alert("Log example", logger.Any("my_key", "my_value"))
	myLogger.Emergency("Log example", logger.Any("my_key", "my_value"))

	output.Contains([]string{
		"lvl: debug | msg: Log example | ctx: <my_key:my_value>",
		"lvl: info | msg: Log example | ctx: <my_key:my_value>",
		"lvl: notice | msg: Log example | ctx: <my_key:my_value>",
		"lvl: warning | msg: Log example | ctx: <my_key:my_value>",
		"lvl: error | msg: Log example | ctx: <my_key:my_value>",
		"lvl: critical | msg: Log example | ctx: <my_key:my_value>",
		"lvl: alert | msg: Log example | ctx: <my_key:my_value>",
		"lvl: emergency | msg: Log example | ctx: <my_key:my_value>",
	})

	//Output:
}

func ExampleLogger_jsonFormatter() {
	output := &Output{}
	myLogger := logger.NewLogger(
		handler.Stream(output, formatter.NewJSONEncoder()),
	)

	myLogger.Debug("Log example", logger.Any("my_key", "my_value"))
	myLogger.Info("Log example", logger.Any("my_key", "my_value"))
	myLogger.Notice("Log example", logger.Any("my_key", "my_value"))
	myLogger.Warning("Log example", logger.Any("my_key", "my_value"))
	myLogger.Error("Log example", logger.Any("my_key", "my_value"))
	myLogger.Critical("Log example", logger.Any("my_key", "my_value"))
	myLogger.Alert("Log example", logger.Any("my_key", "my_value"))
	myLogger.Emergency("Log example", logger.Any("my_key", "my_value"))

	output.Contains([]string{
		`{"Message":"Log example","Level":7,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":6,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":5,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":4,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":3,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":2,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":1,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":0,"Context":{"my_key":"my_value"}}`,
	})

	//Output:
}

func ExampleLogger_minLevelFilterHandler() {
	output := &Output{}
	myLogger := logger.NewLogger(
		middleware.MinLevelFilter(logger.WarningLevel)(
			handler.Stream(output, formatter.NewDefaultFormatter(formatter.WithContext(true))),
		),
	)

	myLogger.Debug("Log example", logger.Any("my_key", "my_value"))
	myLogger.Info("Log example", logger.Any("my_key", "my_value"))
	myLogger.Notice("Log example", logger.Any("my_key", "my_value"))
	myLogger.Warning("Log example", logger.Any("my_key", "my_value"))
	myLogger.Error("Log example", logger.Any("my_key", "my_value"))
	myLogger.Critical("Log example", logger.Any("my_key", "my_value"))
	myLogger.Alert("Log example", logger.Any("my_key", "my_value"))
	myLogger.Emergency("Log example", logger.Any("my_key", "my_value"))

	output.Contains([]string{
		`<warning> Log example {"my_key":"my_value"}`,
		`<error> Log example {"my_key":"my_value"}`,
		`<critical> Log example {"my_key":"my_value"}`,
		`<alert> Log example {"my_key":"my_value"}`,
		`<emergency> Log example {"my_key":"my_value"}`,
	})

	//Output:
}

func ExampleLogger_groupHandler() {
	output := &Output{}
	output2 := &Output{}
	myLogger := logger.NewLogger(
		handler.Group(
			handler.Stream(output, formatter.NewJSONEncoder()),
			handler.Stream(output2, formatter.NewDefaultFormatter(formatter.WithContext(true))),
		),
	)

	myLogger.Debug("Log example", logger.Any("my_key", "my_value"))
	myLogger.Info("Log example", logger.Any("my_key", "my_value"))
	myLogger.Notice("Log example", logger.Any("my_key", "my_value"))
	myLogger.Warning("Log example", logger.Any("my_key", "my_value"))
	myLogger.Error("Log example", logger.Any("my_key", "my_value"))
	myLogger.Critical("Log example", logger.Any("my_key", "my_value"))
	myLogger.Alert("Log example", logger.Any("my_key", "my_value"))
	myLogger.Emergency("Log example", logger.Any("my_key", "my_value"))

	output.Contains([]string{
		`{"Message":"Log example","Level":7,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":6,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":5,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":4,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":3,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":2,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":1,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":0,"Context":{"my_key":"my_value"}}`,
	})

	output2.Contains([]string{
		`<debug> Log example {"my_key":"my_value"}`,
		`<info> Log example {"my_key":"my_value"}`,
		`<notice> Log example {"my_key":"my_value"}`,
		`<warning> Log example {"my_key":"my_value"}`,
		`<error> Log example {"my_key":"my_value"}`,
		`<critical> Log example {"my_key":"my_value"}`,
		`<alert> Log example {"my_key":"my_value"}`,
		`<emergency> Log example {"my_key":"my_value"}`,
	})

	//Output:
}

func ExampleLogger_placeholderMiddleware() {
	output := &Output{}
	myLogger := logger.NewLogger(
		handler.Stream(output, formatter.NewLine("%[2]s %[1]s%.[3]s")),
	)
	myLogger.Wrap(middleware.Placeholder())

	myLogger.Debug("Log %ctx_key% example", logger.Any("ctx_key", false))
	myLogger.Info("Log %ctx_key% example", logger.Any("ctx_key", 1234))
	myLogger.Warning("Log %ctx_key% example", logger.Any("ctx_key", 5*time.Second))
	myLogger.Error("Log %ctx_key% example", logger.Any("ctx_key", "ctx_value"))
	myLogger.Alert("Log %ctx_key% example", logger.Any("ctx_key", struct{ attr string }{attr: "attrValue"}))
	myLogger.Critical("Log %ctx_key% example another value %ctx_key2%", logger.Any("ctx_key", false), logger.Any("ctx_key2", 1234))

	output.Contains([]string{
		`debug Log false example`,
		`info Log 1234 example`,
		`warning Log 5s example`,
		`error Log ctx_value example`,
		`alert Log {attrValue} example`,
		`critical Log false example another value 1234`,
	})

	//Output:
}

func ExampleLogger_wrapHandler() {
	output := &Output{}
	myLogger := logger.NewLogger(
		handler.Stream(output, formatter.NewDefaultFormatter(formatter.WithContext(true))),
	)
	myLogger.Wrap(middleware.MinLevelFilter(logger.WarningLevel))

	myLogger.Debug("Log example", logger.Any("my_key", "my_value"))
	myLogger.Info("Log example", logger.Any("my_key", "my_value"))
	myLogger.Notice("Log example", logger.Any("my_key", "my_value"))
	myLogger.Warning("Log example", logger.Any("my_key", "my_value"))
	myLogger.Error("Log example", logger.Any("my_key", "my_value"))
	myLogger.Critical("Log example", logger.Any("my_key", "my_value"))
	myLogger.Alert("Log example", logger.Any("my_key", "my_value"))
	myLogger.Emergency("Log example", logger.Any("my_key", "my_value"))
	output.Contains([]string{
		`<warning> Log example {"my_key":"my_value"}`,
		`<error> Log example {"my_key":"my_value"}`,
		`<critical> Log example {"my_key":"my_value"}`,
		`<alert> Log example {"my_key":"my_value"}`,
		`<emergency> Log example {"my_key":"my_value"}`,
	})

	//Output:
}

func ExampleLogger_timeRotateHandler() {
	lineFormatter := formatter.NewLine("lvl: %[2]s | msg: %[1]s | ctx: %[3]v")

	rotateLogHandler, _ := handler.TimeRotateFileStream(os.TempDir()+"/%s.log", time.Stamp, lineFormatter, 1*time.Second)
	myLogger := logger.NewLogger(rotateLogHandler)

	myLogger.Debug("Log example", logger.Any("ctx_key", "ctx_value"))
	myLogger.Info("Log example", logger.Any("ctx_key", "ctx_value"))
	myLogger.Warning("Log example", logger.Any("ctx_key", "ctx_value"))
	time.Sleep(1 * time.Second)
	myLogger.Error("Log example", logger.Any("ctx_key", "ctx_value"))
	myLogger.Alert("Log example", logger.Any("ctx_key", "ctx_value"))
	myLogger.Critical("Log example", logger.Any("ctx_key", "ctx_value"))

	//Output:
}

func ExampleLogger_logRotateHandler() {
	lineFormatter := formatter.NewLine("lvl: %[2]s | msg: %[1]s | ctx: %[3]v")

	rotateLogHandler, _ := handler.LogRotateFileStream("test", os.TempDir()+"/%s.log", time.Stamp, lineFormatter, 1*time.Second)
	myLogger := logger.NewLogger(rotateLogHandler)

	myLogger.Debug("Log example", logger.Any("ctx_key", "ctx_value"))
	myLogger.Info("Log example", logger.Any("ctx_key", "ctx_value"))
	myLogger.Warning("Log example", logger.Any("ctx_key", "ctx_value"))
	time.Sleep(1 * time.Second)
	myLogger.Error("Log example", logger.Any("ctx_key", "ctx_value"))
	myLogger.Alert("Log example", logger.Any("ctx_key", "ctx_value"))
	myLogger.Critical("Log example", logger.Any("ctx_key", "ctx_value"))

	//Output:
}

// You can run the command below to show syslog messages
// syslog -F '$Time $Host $(Sender)[$(PID)] <$((Level)(str))>: $Message'
//Apr 26 12:22:06 hades my_go_logger[69302] <Notice>: <notice> Log example2 {"ctx_key":"ctx_value"}
//Apr 26 12:22:06 hades my_go_logger[69302] <Warning>: <warning> Log example3 {"ctx_key":"ctx_value"}
//Apr 26 12:22:06 hades my_go_logger[69302] <Error>: <error> Log example4 {"ctx_key":"ctx_value"}
//Apr 26 12:22:06 hades my_go_logger[69302] <Critical>: <critical> Log example5 {"ctx_key":"ctx_value"}
//Apr 26 12:22:06 hades my_go_logger[69302] <Alert>: <alert> Log example6 {"ctx_key":"ctx_value"}
//Apr 26 12:22:06 hades my_go_logger[69302] <Emergency>: <emergency> Log example7 {"ctx_key":"ctx_value"}
func ExampleLogger_syslogHandler() {
	syslogHandler, _ := handler.Syslog(
		formatter.NewDefaultFormatter(formatter.WithContext(true)),
		"",
		"",
		syslog.LOG_DEBUG,
		"my_go_logger")
	myLogger := logger.NewLogger(syslogHandler)

	myLogger.Debug("Log example", logger.Any("ctx_key", "ctx_value"))
	myLogger.Info("Log example1", logger.Any("ctx_key", "ctx_value"))
	myLogger.Notice("Log example2", logger.Any("ctx_key", "ctx_value"))
	myLogger.Warning("Log example3", logger.Any("ctx_key", "ctx_value"))
	myLogger.Error("Log example4", logger.Any("ctx_key", "ctx_value"))
	myLogger.Critical("Log example5", logger.Any("ctx_key", "ctx_value"))
	myLogger.Alert("Log example6", logger.Any("ctx_key", "ctx_value"))
	myLogger.Emergency("Log example7", logger.Any("ctx_key", "ctx_value"))

	//Output:
}

type Output struct {
	bytes.Buffer
}

func (o *Output) Contains(str []string) {
	b := o.String()
	for _, s := range str {
		if strings.Contains(b, s) != true {
			fmt.Printf("buffer %s must contain %s\n", b, s)
		}
	}
}

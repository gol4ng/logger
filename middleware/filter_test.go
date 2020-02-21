package middleware_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/middleware"
)

func TestFilter_HandleWithoutExclusion(t *testing.T) {
	mockHandler := func(entry logger.Entry) error {
		assert.Equal(t, "my_log_message", entry.Message)
		assert.Equal(t, logger.DebugLevel, entry.Level)
		contextStr := entry.Context.String()
		assert.Contains(t, contextStr, "<my_key:my_overwritten_value>")
		assert.Contains(t, contextStr, "<my_entry_key:my_entry_value>")

		return nil
	}

	filter := middleware.Filter(func(e logger.Entry) bool {
		return true
	})

	logEntry := logger.Entry{
		Message: "my_log_message",
		Level:   logger.DebugLevel,
		Context: nil,
	}

	assert.Nil(t, filter(mockHandler)(logEntry))
}

func TestFilter_HandleWithExclusion(t *testing.T) {
	mockHandler := func(entry logger.Entry) error {
		assert.Fail(t, "Handler function must ne be called")
		return nil
	}

	filter := middleware.Filter(func(e logger.Entry) bool {
		return true
	})

	logEntry := logger.Entry{
		Message: "my_log_message",
		Level:   logger.DebugLevel,
		Context: nil,
	}

	assert.Nil(t, filter(mockHandler)(logEntry))
}

func TestNewMinLevelFilter(t *testing.T) {
	levels := [8]logger.Level{
		logger.DebugLevel,
		logger.InfoLevel,
		logger.NoticeLevel,
		logger.WarningLevel,
		logger.ErrorLevel,
		logger.CriticalLevel,
		logger.AlertLevel,
		logger.EmergencyLevel,
	}

	tests := []struct {
		name             string
		lvl              logger.Level
		logLevelsHandled [8]bool
	}{
		{name: "test min lvl DEBUG", lvl: logger.DebugLevel, logLevelsHandled: [8]bool{true, true, true, true, true, true, true, true}},
		{name: "test min lvl INFO", lvl: logger.InfoLevel, logLevelsHandled: [8]bool{false, true, true, true, true, true, true, true}},
		{name: "test min lvl NOTICE", lvl: logger.NoticeLevel, logLevelsHandled: [8]bool{false, false, true, true, true, true, true, true}},
		{name: "test min lvl WARNING", lvl: logger.WarningLevel, logLevelsHandled: [8]bool{false, false, false, true, true, true, true, true}},
		{name: "test min lvl ERROR", lvl: logger.ErrorLevel, logLevelsHandled: [8]bool{false, false, false, false, true, true, true, true}},
		{name: "test min lvl CRITICAL", lvl: logger.CriticalLevel, logLevelsHandled: [8]bool{false, false, false, false, false, true, true, true}},
		{name: "test min lvl ALERT", lvl: logger.AlertLevel, logLevelsHandled: [8]bool{false, false, false, false, false, false, true, true}},
		{name: "test min lvl EMERGENCY", lvl: logger.EmergencyLevel, logLevelsHandled: [8]bool{false, false, false, false, false, false, false, true}},
	}
	for _, tt := range tests {
		for i, logLevel := range levels {
			t.Run(tt.name, func(t *testing.T) {
				entry := logger.Entry{Level: logLevel}
				mockCalled := false
				mockHandler := func(entry logger.Entry) error {
					mockCalled = true
					return nil
				}

				filter := middleware.MinLevelFilter(tt.lvl)

				assert.Nil(t, filter(mockHandler)(entry))
				assert.Equal(t, tt.logLevelsHandled[i], mockCalled)
			})
		}
	}
}

func TestNewMaxLevelFilter(t *testing.T) {
	levels := [8]logger.Level{
		logger.DebugLevel,
		logger.InfoLevel,
		logger.NoticeLevel,
		logger.WarningLevel,
		logger.ErrorLevel,
		logger.CriticalLevel,
		logger.AlertLevel,
		logger.EmergencyLevel,
	}

	tests := []struct {
		name             string
		lvl              logger.Level
		logLevelsHandled [8]bool
	}{
		{name: "test max lvl DEBUG", lvl: logger.DebugLevel, logLevelsHandled: [8]bool{true, false, false, false, false, false, false, false}},
		{name: "test max lvl INFO", lvl: logger.InfoLevel, logLevelsHandled: [8]bool{true, true, false, false, false, false, false, false}},
		{name: "test max lvl NOTICE", lvl: logger.NoticeLevel, logLevelsHandled: [8]bool{true, true, true, false, false, false, false, false}},
		{name: "test max lvl WARNING", lvl: logger.WarningLevel, logLevelsHandled: [8]bool{true, true, true, true, false, false, false, false}},
		{name: "test max lvl ERROR", lvl: logger.ErrorLevel, logLevelsHandled: [8]bool{true, true, true, true, true, false, false, false}},
		{name: "test max lvl CRITICAL", lvl: logger.CriticalLevel, logLevelsHandled: [8]bool{true, true, true, true, true, true, false, false}},
		{name: "test max lvl ALERT", lvl: logger.AlertLevel, logLevelsHandled: [8]bool{true, true, true, true, true, true, true, false}},
		{name: "test max lvl EMERGENCY", lvl: logger.EmergencyLevel, logLevelsHandled: [8]bool{true, true, true, true, true, true, true, true}},
	}
	for _, tt := range tests {
		for i, logLevel := range levels {
			t.Run(tt.name, func(t *testing.T) {
				entry := logger.Entry{Level: logLevel}
				mockCalled := false
				mockHandler := func(entry logger.Entry) error {
					mockCalled = true
					return nil
				}

				filter := middleware.MaxLevelFilter(tt.lvl)

				assert.Nil(t, filter(mockHandler)(entry))
				assert.Equal(t, tt.logLevelsHandled[i], mockCalled)
			})
		}
	}
}

func TestNewRangeLevelFilter(t *testing.T) {
	levels := [8]logger.Level{
		logger.DebugLevel,
		logger.InfoLevel,
		logger.NoticeLevel,
		logger.WarningLevel,
		logger.ErrorLevel,
		logger.CriticalLevel,
		logger.AlertLevel,
		logger.EmergencyLevel,
	}

	tests := []struct {
		name             string
		level1           logger.Level
		level2           logger.Level
		logLevelsHandled [8]bool
	}{
		{name: "test between DEBUG/EMERGENCY with log level %s", level1: logger.DebugLevel, level2: logger.EmergencyLevel, logLevelsHandled: [8]bool{true, true, true, true, true, true, true, true}},
		{name: "test between INFO/ALERT with log level %s", level1: logger.InfoLevel, level2: logger.AlertLevel, logLevelsHandled: [8]bool{false, true, true, true, true, true, true, false}},
		{name: "test between EMERGENCY/DEBUG with log level %s", level1: logger.EmergencyLevel, level2: logger.DebugLevel, logLevelsHandled: [8]bool{true, true, true, true, true, true, true, true}},
	}
	for _, tt := range tests {
		for i, logLevel := range levels {
			t.Run(fmt.Sprintf(tt.name, logLevel), func(t *testing.T) {
				entry := logger.Entry{Level: logLevel}
				mockCalled := false
				mockHandler := func(entry logger.Entry) error {
					mockCalled = true
					return nil
				}

				filter := middleware.RangeLevelFilter(tt.level1, tt.level2)

				assert.Nil(t, filter(mockHandler)(entry))
				assert.Equal(t, tt.logLevelsHandled[i], mockCalled)
			})
		}
	}
}


// =====================================================================================================================
// ================================================= EXAMPLES ==========================================================
// =====================================================================================================================

func ExampleMinLevelFilter() {
	streamHandler := handler.Stream(os.Stdout, formatter.NewDefaultFormatter())

	minLvlFilterHandler := middleware.MinLevelFilter(logger.WarningLevel)(streamHandler)
	_ = minLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.DebugLevel})
	_ = minLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.InfoLevel})
	_ = minLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.NoticeLevel})
	_ = minLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.WarningLevel})
	_ = minLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.ErrorLevel})
	_ = minLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.CriticalLevel})
	_ = minLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.AlertLevel})
	_ = minLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.EmergencyLevel})

	//Output:
	//<warning> will be printed
	//<error> will be printed
	//<critical> will be printed
	//<alert> will be printed
	//<emergency> will be printed
}

func ExampleRangeLevelFilter() {
	streamHandler := handler.Stream(os.Stdout, formatter.NewDefaultFormatter())

	rangeLvlFilterHandler := middleware.RangeLevelFilter(logger.InfoLevel, logger.WarningLevel)(streamHandler)

	_ = rangeLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.DebugLevel})
	_ = rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.InfoLevel})
	_ = rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.NoticeLevel})
	_ = rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.WarningLevel})
	_ = rangeLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.ErrorLevel})
	_ = rangeLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.CriticalLevel})
	_ = rangeLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.AlertLevel})
	_ = rangeLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.EmergencyLevel})

	//Output:
	//<info> will be printed
	//<notice> will be printed
	//<warning> will be printed
}

func ExampleCustomFilter() {
	streamHandler := handler.Stream(os.Stdout, formatter.NewDefaultFormatter())

	rangeLvlFilterHandler := middleware.Filter(func(e logger.Entry) bool {
		return e.Level == logger.InfoLevel || e.Level == logger.AlertLevel
	})(streamHandler)

	_ = rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.DebugLevel})
	_ = rangeLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.InfoLevel})
	_ = rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.NoticeLevel})
	_ = rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.WarningLevel})
	_ = rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.ErrorLevel})
	_ = rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.CriticalLevel})
	_ = rangeLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.AlertLevel})
	_ = rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.EmergencyLevel})

	//Output:
	//<debug> will be printed
	//<notice> will be printed
	//<warning> will be printed
	//<error> will be printed
	//<critical> will be printed
	//<emergency> will be printed
}

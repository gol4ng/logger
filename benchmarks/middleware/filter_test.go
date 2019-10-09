package middleware_test

import (
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/middleware"
)

func BenchmarkMinLevelFilterHandler(b *testing.B) {
	b.ReportAllocs()

	minLvlFilterHandler := middleware.MinLevelFilter(logger.InfoLevel)(logger.NopHandler)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = minLvlFilterHandler(logger.Entry{Message: "This log message is really logged.", Level: logger.InfoLevel})
		_ = minLvlFilterHandler(logger.Entry{Message: "This log message will be excluded.", Level: logger.DebugLevel})
	}
}

func BenchmarkRangeLevelFilterHandler(b *testing.B) {
	b.ReportAllocs()

	rangeLevelHandler := middleware.RangeLevelFilter(logger.InfoLevel, logger.WarningLevel)(logger.NopHandler)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = rangeLevelHandler(logger.Entry{Message: "This log message is really logged.", Level: logger.InfoLevel})
		_ = rangeLevelHandler(logger.Entry{Message: "This log message is really logged.", Level: logger.WarningLevel})
		_ = rangeLevelHandler(logger.Entry{Message: "This log message will be excluded.", Level: logger.DebugLevel})
		_ = rangeLevelHandler(logger.Entry{Message: "This log message will be excluded.", Level: logger.ErrorLevel})
	}
}

package benchmark_handler_test

import (
	"testing"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/handler"
)

func BenchmarkMinLevelFilterHandler(b *testing.B) {
	b.ReportAllocs()

	myLogger := logger.NewLogger(handler.NewMinLevelFilter(&logger.NilHandler{}, logger.InfoLevel))

	for n := 0; n < b.N; n++ {
		myLogger.Info("This log message is really logged.", nil)
		myLogger.Debug("This log message will be excluded.", nil)
	}
}

func BenchmarkRangeLevelFilterHandler(b *testing.B) {
	b.ReportAllocs()

	myLogger := logger.NewLogger(handler.NewRangeLevelFilter(&logger.NilHandler{}, logger.InfoLevel, logger.WarnLevel))

	for n := 0; n < b.N; n++ {
		myLogger.Info("This log message is really logged.", nil)
		myLogger.Warn("This log message is really logged.", nil)
		myLogger.Debug("This log message will be excluded.", nil)
		myLogger.Error("This log message will be excluded.", nil)
	}
}

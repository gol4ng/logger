package benchmark_handler_test

import (
	"testing"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/handler"
)

func BenchmarkMinLevelFilterHandler(b *testing.B) {
	b.ReportAllocs()

	minLvlFilterHandler := handler.NewMinLevelFilter(&logger.NilHandler{}, logger.InfoLevel)

	for n := 0; n < b.N; n++ {
		minLvlFilterHandler.Handle(logger.Entry{Message: "This log message is really logged.", Level: logger.InfoLevel})
		minLvlFilterHandler.Handle(logger.Entry{Message: "This log message will be excluded.", Level: logger.DebugLevel})
	}
}

func BenchmarkRangeLevelFilterHandler(b *testing.B) {
	b.ReportAllocs()

	rangeLevelHandler := handler.NewRangeLevelFilter(&logger.NilHandler{}, logger.InfoLevel, logger.WarnLevel)

	for n := 0; n < b.N; n++ {
		rangeLevelHandler.Handle(logger.Entry{Message: "This log message is really logged.", Level: logger.InfoLevel})
		rangeLevelHandler.Handle(logger.Entry{Message: "This log message is really logged.", Level: logger.WarnLevel})
		rangeLevelHandler.Handle(logger.Entry{Message: "This log message will be excluded.", Level: logger.DebugLevel})
		rangeLevelHandler.Handle(logger.Entry{Message: "This log message will be excluded.", Level: logger.ErrorLevel})
	}
}

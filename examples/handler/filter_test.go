// +build example

package example_handler_test

import (
	"os"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/handler"
)

func ExampleMinLevelFilterHandler() {
	streamHandler := handler.NewStream(os.Stdout, logger.NewNilFormatter())

	minLvlFilterHandler := handler.NewMinLevelFilter(streamHandler, logger.WarnLevel)

	myLogger := logger.Logger{HandlerInterface: minLvlFilterHandler}

	myLogger.Debug("will be excluded", nil)
	myLogger.Info("will be excluded", nil)
	myLogger.Warn("will be printed", nil)
	myLogger.Error("will be printed", nil)
	myLogger.Panic("will be printed", nil)
	myLogger.Fatal("will be printed", nil)

	//Output:
	// {will be printed warn <nil>}
	// {will be printed error <nil>}
	// {will be printed panic <nil>}
	// {will be printed fatal <nil>}
}

func ExampleRangeLevelFilterHandler() {
	streamHandler := handler.NewStream(os.Stdout, logger.NewNilFormatter())

	rangeLvlFilterHandler := handler.NewRangeLevelFilter(streamHandler, logger.InfoLevel, logger.WarnLevel)

	myLogger := logger.Logger{HandlerInterface: rangeLvlFilterHandler}

	myLogger.Debug("will be excluded", nil)
	myLogger.Info("will be printed", nil)
	myLogger.Warn("will be printed", nil)
	myLogger.Error("will be excluded", nil)
	myLogger.Panic("will be excluded", nil)
	myLogger.Fatal("will be excluded", nil)

	//Output:
	// {will be printed info <nil>}
	// {will be printed warn <nil>}
}

func ExampleCustomFilterHandler() {
	streamHandler := handler.NewStream(os.Stdout, logger.NewNilFormatter())

	rangeLvlFilterHandler := handler.NewFilter(streamHandler, func(e logger.Entry) bool {
		return e.Level == logger.InfoLevel || e.Level == logger.PanicLevel
	})

	myLogger := logger.Logger{HandlerInterface: rangeLvlFilterHandler}

	myLogger.Debug("will be printed", nil)
	myLogger.Info("will be excluded", nil)
	myLogger.Warn("will be printed", nil)
	myLogger.Error("will be printed", nil)
	myLogger.Panic("will be excluded", nil)
	myLogger.Fatal("will be printed", nil)

	//Output:
	// {will be printed debug <nil>}
	// {will be printed warn <nil>}
	// {will be printed error <nil>}
	// {will be printed fatal <nil>}
}

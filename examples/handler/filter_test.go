package example_handler_test

import (
	"os"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/handler"
)

func ExampleMinLevelFilterHandler() {
	myString := "Log example"

	streamHandler := handler.NewStream(os.Stdout, logger.NewNilFormatter())

	minLvlFilterHandler := handler.NewMinLevelFilter(streamHandler, logger.WarnLevel)

	myLogger := logger.Logger{HandlerInterface: minLvlFilterHandler}

	myLogger.Debug(myString, nil)
	myLogger.Info(myString, nil)
	myLogger.Warn(myString, nil)
	myLogger.Error(myString, nil)

	//Output:
	// {Log example warn <nil>}
	// {Log example error <nil>}
}

func ExampleRangeLevelFilterHandler() {
	myString := "Log example"

	streamHandler := handler.NewStream(os.Stdout, logger.NewNilFormatter())

	minLvlFilterHandler := handler.NewRangeLevelFilter(streamHandler, logger.InfoLevel, logger.WarnLevel)

	myLogger := logger.Logger{HandlerInterface: minLvlFilterHandler}

	myLogger.Debug(myString, nil)
	myLogger.Info(myString, nil)
	myLogger.Warn(myString, nil)
	myLogger.Error(myString, nil)

	//Output:
	// {Log example info <nil>}
	// {Log example warn <nil>}
}

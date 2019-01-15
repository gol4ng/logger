package example_handler_test

import (
	"os"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/formatter"
	"github.com/instabledesign/logger/handler"
)

func ExampleLoggerLineFormatter() {
	lineLogHandler := handler.NewStream(os.Stdout, formatter.NewLine("%[2]s | %[1]s"))
	myLogger := logger.NewLogger(lineLogHandler)

	myLogger.Info("Log example", nil)

	//Output:
	// info | Log example
}

func ExampleLoggerJsonFormatter() {
	lineLogHandler := handler.NewStream(os.Stdout, formatter.NewJson())
	myLogger := logger.NewLogger(lineLogHandler)

	myLogger.Info("Log example", nil)

	//Output:
	// {"Message":"Log example","Level":0,"Context":null}
}

func ExampleLoggerMinLevelFilterHandler() {
	lineLogHandler := handler.NewStream(os.Stdout, &logger.NilFormatter{})
	filterLogHandler := handler.NewMinLevelFilter(lineLogHandler, logger.WarnLevel)
	myLogger := logger.NewLogger(filterLogHandler)

	myLogger.Debug("Log example", nil)
	myLogger.Info("Log example", nil)
	myLogger.Warn("Log example", nil)
	myLogger.Error("Log example", nil)
	myLogger.Panic("Log example", nil)
	myLogger.Fatal("Log example", nil)

	//Output:
	// {Log example warn <nil>}
	// {Log example error <nil>}
	// {Log example panic <nil>}
	// {Log example fatal <nil>}
}

func ExampleLoggerGroupHandler() {
	jsonLogHandler := handler.NewStream(os.Stdout, &formatter.Json{})
	lineLogHandler := handler.NewStream(os.Stdout, &logger.NilFormatter{})
	groupLogHandler := handler.NewGroup(jsonLogHandler, lineLogHandler)
	myLogger := logger.NewLogger(groupLogHandler)

	myLogger.Debug("Log example", nil)
	myLogger.Info("Log example", nil)
	myLogger.Warn("Log example", nil)
	myLogger.Error("Log example", nil)
	myLogger.Panic("Log example", nil)
	myLogger.Fatal("Log example", nil)

	//Output:
	// {"Message":"Log example","Level":-1,"Context":null}
	// {Log example debug <nil>}
	// {"Message":"Log example","Level":0,"Context":null}
	// {Log example info <nil>}
	// {"Message":"Log example","Level":1,"Context":null}
	// {Log example warn <nil>}
	// {"Message":"Log example","Level":2,"Context":null}
	// {Log example error <nil>}
	// {"Message":"Log example","Level":3,"Context":null}
	// {Log example panic <nil>}
	// {"Message":"Log example","Level":4,"Context":null}
	// {Log example fatal <nil>}
}

func ExampleLoggerWrapHandler() {
	lineLogHandler := handler.NewStream(os.Stdout, &logger.NilFormatter{})
	myLogger := logger.NewLogger(lineLogHandler)
	myLogger.Wrap(handler.NewMinLevelWrapper(logger.WarnLevel))

	myLogger.Debug("Log example", nil)
	myLogger.Info("Log example", nil)
	myLogger.Warn("Log example", nil)
	myLogger.Error("Log example", nil)
	myLogger.Panic("Log example", nil)
	myLogger.Fatal("Log example", nil)

	//Output:
	// {Log example warn <nil>}
	// {Log example error <nil>}
	// {Log example panic <nil>}
	// {Log example fatal <nil>}
}

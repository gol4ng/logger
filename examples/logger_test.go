package example_handler_test

import (
	"os"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/formatter"
	"github.com/instabledesign/logger/handler"
)

func ExampleLoggerLineFormatter() {
	lineLogHandler := handler.NewStream(os.Stdout, formatter.NewLine("lvl: %[2]s | msg: %[1]s | ctx: %[3]v"))
	myLogger := logger.NewLogger(lineLogHandler)

	myLogger.Info("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})

	//Output:
	// lvl: info | msg: Log example | ctx: &map[ctx_key:ctx_value]
}

func ExampleLoggerJsonFormatter() {
	lineLogHandler := handler.NewStream(os.Stdout, formatter.NewJson())
	myLogger := logger.NewLogger(lineLogHandler)

	myLogger.Info("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})

	//Output:
	// {"Message":"Log example","Level":0,"Context":{"ctx_key":"ctx_value"}}
}

func ExampleLoggerMinLevelFilterHandler() {
	lineLogHandler := handler.NewStream(os.Stdout, formatter.NewLine("%s %s %s"))
	filterLogHandler := handler.NewMinLevelFilter(lineLogHandler, logger.WarnLevel)
	myLogger := logger.NewLogger(filterLogHandler)

	myLogger.Debug("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Info("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Warn("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Error("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Panic("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Fatal("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})

	//Output:
	// Log example warn &map[ctx_key:ctx_value]
	// Log example error &map[ctx_key:ctx_value]
	// Log example panic &map[ctx_key:ctx_value]
	// Log example fatal &map[ctx_key:ctx_value]
}

func ExampleLoggerGroupHandler() {
	jsonLogHandler := handler.NewStream(os.Stdout, &formatter.Json{})
	lineLogHandler := handler.NewStream(os.Stdout, formatter.NewLine("%s %s %s"))
	groupLogHandler := handler.NewGroup(jsonLogHandler, lineLogHandler)
	myLogger := logger.NewLogger(groupLogHandler)

	myLogger.Debug("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Info("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Warn("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Error("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Panic("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Fatal("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})

	//Output:
	// {"Message":"Log example","Level":-1,"Context":{"ctx_key":"ctx_value"}}
	// Log example debug &map[ctx_key:ctx_value]
	// {"Message":"Log example","Level":0,"Context":{"ctx_key":"ctx_value"}}
	// Log example info &map[ctx_key:ctx_value]
	// {"Message":"Log example","Level":1,"Context":{"ctx_key":"ctx_value"}}
	// Log example warn &map[ctx_key:ctx_value]
	// {"Message":"Log example","Level":2,"Context":{"ctx_key":"ctx_value"}}
	// Log example error &map[ctx_key:ctx_value]
	// {"Message":"Log example","Level":3,"Context":{"ctx_key":"ctx_value"}}
	// Log example panic &map[ctx_key:ctx_value]
	// {"Message":"Log example","Level":4,"Context":{"ctx_key":"ctx_value"}}
	// Log example fatal &map[ctx_key:ctx_value]
}

func ExampleLoggerWrapHandler() {
	lineLogHandler := handler.NewStream(os.Stdout, formatter.NewLine("%s %s %s"))
	myLogger := logger.NewLogger(lineLogHandler)
	myLogger.Wrap(handler.NewMinLevelWrapper(logger.WarnLevel))

	myLogger.Debug("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Info("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Warn("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Error("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Panic("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Fatal("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})

	//Output:
	// Log example warn &map[ctx_key:ctx_value]
	// Log example error &map[ctx_key:ctx_value]
	// Log example panic &map[ctx_key:ctx_value]
	// Log example fatal &map[ctx_key:ctx_value]
}

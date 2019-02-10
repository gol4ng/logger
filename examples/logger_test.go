package example_handler_test

import (
	"log/syslog"
	"os"
	"time"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
)

func ExampleLoggerLineFormatter() {
	lineLogHandler := handler.NewStream(os.Stdout, formatter.NewLine("lvl: %[2]s | msg: %[1]s | ctx: %[3]v"))
	myLogger := logger.NewLogger(lineLogHandler)

	myLogger.Debug("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Info("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Notice("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Warning("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Error("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Critical("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Alert("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Emergency("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})

	//Output:
	//lvl: debug | msg: Log example | ctx: &map[ctx_key:ctx_value]
	//lvl: info | msg: Log example | ctx: &map[ctx_key:ctx_value]
	//lvl: notice | msg: Log example | ctx: &map[ctx_key:ctx_value]
	//lvl: warning | msg: Log example | ctx: &map[ctx_key:ctx_value]
	//lvl: error | msg: Log example | ctx: &map[ctx_key:ctx_value]
	//lvl: critical | msg: Log example | ctx: &map[ctx_key:ctx_value]
	//lvl: alert | msg: Log example | ctx: &map[ctx_key:ctx_value]
	//lvl: emergency | msg: Log example | ctx: &map[ctx_key:ctx_value]
}

func ExampleLoggerJsonFormatter() {
	lineLogHandler := handler.NewStream(os.Stdout, formatter.NewJson())
	myLogger := logger.NewLogger(lineLogHandler)

	myLogger.Debug("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Info("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Notice("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Warning("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Error("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Critical("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Alert("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Emergency("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})

	//Output:
	//{"Message":"Log example","Level":7,"Context":{"ctx_key":"ctx_value"}}
	//{"Message":"Log example","Level":6,"Context":{"ctx_key":"ctx_value"}}
	//{"Message":"Log example","Level":5,"Context":{"ctx_key":"ctx_value"}}
	//{"Message":"Log example","Level":4,"Context":{"ctx_key":"ctx_value"}}
	//{"Message":"Log example","Level":3,"Context":{"ctx_key":"ctx_value"}}
	//{"Message":"Log example","Level":2,"Context":{"ctx_key":"ctx_value"}}
	//{"Message":"Log example","Level":1,"Context":{"ctx_key":"ctx_value"}}
	//{"Message":"Log example","Level":0,"Context":{"ctx_key":"ctx_value"}}
}

func ExampleLoggerMinLevelFilterHandler() {
	lineLogHandler := handler.NewStream(os.Stdout, formatter.NewDefaultFormatter())
	filterLogHandler := handler.NewMinLevelFilter(lineLogHandler, logger.WarningLevel)
	myLogger := logger.NewLogger(filterLogHandler)

	myLogger.Debug("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Info("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Notice("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Warning("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Error("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Critical("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Alert("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Emergency("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})

	//Output:
	//warning Log example
	//error Log example
	//critical Log example
	//alert Log example
	//emergency Log example
}

func ExampleLoggerGroupHandler() {
	jsonLogHandler := handler.NewStream(os.Stdout, &formatter.Json{})
	lineLogHandler := handler.NewStream(os.Stdout, formatter.NewDefaultFormatter())
	groupLogHandler := handler.NewGroup(jsonLogHandler, lineLogHandler)
	myLogger := logger.NewLogger(groupLogHandler)

	myLogger.Debug("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Info("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Notice("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Warning("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Error("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Critical("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Alert("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Emergency("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})

	//Output:
	//{"Message":"Log example","Level":7,"Context":{"ctx_key":"ctx_value"}}
	//debug Log example
	//{"Message":"Log example","Level":6,"Context":{"ctx_key":"ctx_value"}}
	//info Log example
	//{"Message":"Log example","Level":5,"Context":{"ctx_key":"ctx_value"}}
	//notice Log example
	//{"Message":"Log example","Level":4,"Context":{"ctx_key":"ctx_value"}}
	//warning Log example
	//{"Message":"Log example","Level":3,"Context":{"ctx_key":"ctx_value"}}
	//error Log example
	//{"Message":"Log example","Level":2,"Context":{"ctx_key":"ctx_value"}}
	//critical Log example
	//{"Message":"Log example","Level":1,"Context":{"ctx_key":"ctx_value"}}
	//alert Log example
	//{"Message":"Log example","Level":0,"Context":{"ctx_key":"ctx_value"}}
	//emergency Log example
}

func ExampleLoggerWrapHandler() {
	lineLogHandler := handler.NewStream(os.Stdout, formatter.NewDefaultFormatter())
	myLogger := logger.NewLogger(lineLogHandler)
	myLogger.Wrap(handler.NewMinLevelWrapper(logger.WarningLevel))

	myLogger.Debug("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Info("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Notice("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Warning("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Error("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Critical("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Alert("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Emergency("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})

	//Output:
	//warning Log example
	//error Log example
	//critical Log example
	//alert Log example
	//emergency Log example
}

func ExampleLoggerTimeRotateHandler() {
	lineFormatter := formatter.NewLine("lvl: %[2]s | msg: %[1]s | ctx: %[3]v")

	rotateLogHandler, _ := handler.NewTimeRotateFileStream(os.TempDir()+"%s.log", time.Stamp, lineFormatter, 1*time.Second)
	myLogger := logger.NewLogger(rotateLogHandler)

	myLogger.Debug("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Info("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Warning("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	time.Sleep(1 * time.Second)
	myLogger.Error("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Alert("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Critical("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})

	//Output:
	//
}

func ExampleLoggerLogRotateHandler() {
	lineFormatter := formatter.NewLine("lvl: %[2]s | msg: %[1]s | ctx: %[3]v")

	rotateLogHandler, _ := handler.NewLogRotateFileStream("test", os.TempDir()+"%s.log", time.Stamp, lineFormatter, 1*time.Second)
	myLogger := logger.NewLogger(rotateLogHandler)

	myLogger.Debug("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Info("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Warning("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	time.Sleep(1 * time.Second)
	myLogger.Error("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Alert("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Critical("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})

	//Output:
	//
}

// You can run the command below to show syslog messages
// syslog -F '$Time $Host $(Sender)[$(PID)] <$((Level)(str))>: $Message'
//Jan 22 22:42:14 hades my_go_logger[113] <Notice>: notice Log example2 &map[ctx_key:ctx_value]
//Jan 22 22:42:14 hades my_go_logger[113] <Warning>: warning Log example3 &map[ctx_key:ctx_value]
//Jan 22 22:42:14 hades my_go_logger[113] <Error>: error Log example4 &map[ctx_key:ctx_value]
//Jan 22 22:42:14 hades my_go_logger[113] <Critical>: critical Log example5 &map[ctx_key:ctx_value]
//Jan 22 22:42:14 hades my_go_logger[113] <Alert>: alert Log example6 &map[ctx_key:ctx_value]
//Jan 22 22:42:14 hades my_go_logger[113] <Emergency>: emergency Log example7 &map[ctx_key:ctx_value]
func ExampleLoggerSyslogHandler() {
	syslogHandler, _ := handler.NewSyslog(
		formatter.NewDefaultFormatter(),
		"",
		"",
		syslog.LOG_DEBUG,
		"my_go_logger")
	myLogger := logger.NewLogger(syslogHandler)

	myLogger.Debug("Log example", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Info("Log example1", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Notice("Log example2", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Warning("Log example3", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Error("Log example4", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Critical("Log example5", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Alert("Log example6", &map[string]interface{}{"ctx_key": "ctx_value"})
	myLogger.Emergency("Log example7", &map[string]interface{}{"ctx_key": "ctx_value"})

	//Output:
}

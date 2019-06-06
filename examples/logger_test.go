package example_handler_test

import (
	"bytes"
	"fmt"
	"log/syslog"
	"os"
	"strings"
	"time"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/middleware"
)

func ExampleLoggerCallerHandler() {
	output := &Output{}
	myLogger := logger.NewLogger(
		middleware.Caller(3)(
			handler.Stream(output, formatter.NewLine("lvl: %[2]s | msg: %[1]s | ctx: %[3]v")),
		),
	)

	myLogger.Debug("Log example", nil)
	myLogger.Info("Log example", nil)
	myLogger.Notice("Log example", nil)
	myLogger.Warning("Log example", nil)
	myLogger.Error("Log example", nil)
	myLogger.Critical("Log example", nil)
	myLogger.Alert("Log example", nil)
	myLogger.Emergency("Log example", nil)

	output.Constains([]string{
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

func ExampleLoggerLineFormatter() {
	output := &Output{}
	myLogger := logger.NewLogger(
		handler.Stream(output, formatter.NewLine("lvl: %[2]s | msg: %[1]s | ctx: %[3]v")),
	)

	myLogger.Debug("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Info("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Notice("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Warning("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Error("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Critical("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Alert("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Emergency("Log example", logger.Ctx("my_key", "my_value"))

	output.Constains([]string{
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

func ExampleLoggerJsonFormatter() {
	output := &Output{}
	myLogger := logger.NewLogger(
		handler.Stream(output, formatter.NewJsonEncoder()),
	)

	myLogger.Debug("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Info("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Notice("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Warning("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Error("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Critical("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Alert("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Emergency("Log example", logger.Ctx("my_key", "my_value"))

	output.Constains([]string{
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

func ExampleLoggerMinLevelFilterHandler() {
	output := &Output{}
	myLogger := logger.NewLogger(
		middleware.MinLevelFilter(logger.WarningLevel)(
			handler.Stream(output, formatter.NewDefaultFormatter()),
		),
	)

	myLogger.Debug("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Info("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Notice("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Warning("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Error("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Critical("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Alert("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Emergency("Log example", logger.Ctx("my_key", "my_value"))

	output.Constains([]string{
		`<warning> Log example {"my_key":"my_value"}`,
		`<error> Log example {"my_key":"my_value"}`,
		`<critical> Log example {"my_key":"my_value"}`,
		`<alert> Log example {"my_key":"my_value"}`,
		`<emergency> Log example {"my_key":"my_value"}`,
	})

	//Output:
}

func ExampleLoggerGroupHandler() {
	output := &Output{}
	output2 := &Output{}
	myLogger := logger.NewLogger(
		handler.Group(
			handler.Stream(output, formatter.NewJsonEncoder()),
			handler.Stream(output2, formatter.NewDefaultFormatter()),
		),
	)

	myLogger.Debug("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Info("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Notice("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Warning("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Error("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Critical("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Alert("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Emergency("Log example", logger.Ctx("my_key", "my_value"))

	output.Constains([]string{
		`{"Message":"Log example","Level":7,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":6,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":5,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":4,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":3,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":2,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":1,"Context":{"my_key":"my_value"}}`,
		`{"Message":"Log example","Level":0,"Context":{"my_key":"my_value"}}`,
	})

	output2.Constains([]string{
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

func ExampleLoggerWrapHandler() {
	output := &Output{}
	myLogger := logger.NewLogger(
		handler.Stream(output, formatter.NewDefaultFormatter()),
	)
	myLogger.Wrap(middleware.MinLevelFilter(logger.WarningLevel))

	myLogger.Debug("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Info("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Notice("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Warning("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Error("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Critical("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Alert("Log example", logger.Ctx("my_key", "my_value"))
	myLogger.Emergency("Log example", logger.Ctx("my_key", "my_value"))
	output.Constains([]string{
		`<warning> Log example {"my_key":"my_value"}`,
		`<error> Log example {"my_key":"my_value"}`,
		`<critical> Log example {"my_key":"my_value"}`,
		`<alert> Log example {"my_key":"my_value"}`,
		`<emergency> Log example {"my_key":"my_value"}`,
	})

	//Output:
}

func ExampleLoggerTimeRotateHandler() {
	lineFormatter := formatter.NewLine("lvl: %[2]s | msg: %[1]s | ctx: %[3]v")

	rotateLogHandler, _ := handler.TimeRotateFileStream(os.TempDir()+"%s.log", time.Stamp, lineFormatter, 1*time.Second)
	myLogger := logger.NewLogger(rotateLogHandler)

	myLogger.Debug("Log example", logger.Ctx("ctx_key", "ctx_value"))
	myLogger.Info("Log example", logger.Ctx("ctx_key", "ctx_value"))
	myLogger.Warning("Log example", logger.Ctx("ctx_key", "ctx_value"))
	time.Sleep(1 * time.Second)
	myLogger.Error("Log example", logger.Ctx("ctx_key", "ctx_value"))
	myLogger.Alert("Log example", logger.Ctx("ctx_key", "ctx_value"))
	myLogger.Critical("Log example", logger.Ctx("ctx_key", "ctx_value"))

	//Output:
}

func ExampleLoggerLogRotateHandler() {
	lineFormatter := formatter.NewLine("lvl: %[2]s | msg: %[1]s | ctx: %[3]v")

	rotateLogHandler, _ := handler.LogRotateFileStream("test", os.TempDir()+"%s.log", time.Stamp, lineFormatter, 1*time.Second)
	myLogger := logger.NewLogger(rotateLogHandler)

	myLogger.Debug("Log example", logger.Ctx("ctx_key", "ctx_value"))
	myLogger.Info("Log example", logger.Ctx("ctx_key", "ctx_value"))
	myLogger.Warning("Log example", logger.Ctx("ctx_key", "ctx_value"))
	time.Sleep(1 * time.Second)
	myLogger.Error("Log example", logger.Ctx("ctx_key", "ctx_value"))
	myLogger.Alert("Log example", logger.Ctx("ctx_key", "ctx_value"))
	myLogger.Critical("Log example", logger.Ctx("ctx_key", "ctx_value"))

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
func ExampleLoggerSyslogHandler() {
	syslogHandler, _ := handler.Syslog(
		formatter.NewDefaultFormatter(),
		"",
		"",
		syslog.LOG_DEBUG,
		"my_go_logger")
	myLogger := logger.NewLogger(syslogHandler)

	myLogger.Debug("Log example", logger.Ctx("ctx_key", "ctx_value"))
	myLogger.Info("Log example1", logger.Ctx("ctx_key", "ctx_value"))
	myLogger.Notice("Log example2", logger.Ctx("ctx_key", "ctx_value"))
	myLogger.Warning("Log example3", logger.Ctx("ctx_key", "ctx_value"))
	myLogger.Error("Log example4", logger.Ctx("ctx_key", "ctx_value"))
	myLogger.Critical("Log example5", logger.Ctx("ctx_key", "ctx_value"))
	myLogger.Alert("Log example6", logger.Ctx("ctx_key", "ctx_value"))
	myLogger.Emergency("Log example7", logger.Ctx("ctx_key", "ctx_value"))

	//Output:
}

type Output struct {
	bytes.Buffer
}

func (o *Output) Constains(str []string) {
	b := o.String()
	for _, s := range str {
		if strings.Contains(b, s) != true {
			fmt.Printf("buffer %s must contain %s\n", b, s)
		}
	}
}

package example_handler_test

import (
	"log/syslog"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
)

// You can run the command below to show syslog messages
// syslog -F '$Time $Host $(Sender)[$(PID)] <$((Level)(str))>: $Message'
//Jan 22 22:42:14 hades my_go_logger[113] <Notice>: notice Log example2 &map[ctx_key:ctx_value]
//Jan 22 22:42:14 hades my_go_logger[113] <Warning>: warning Log example3 &map[ctx_key:ctx_value]
//Jan 22 22:42:14 hades my_go_logger[113] <Error>: error Log example4 &map[ctx_key:ctx_value]
//Jan 22 22:42:14 hades my_go_logger[113] <Critical>: critical Log example5 &map[ctx_key:ctx_value]
//Jan 22 22:42:14 hades my_go_logger[113] <Alert>: alert Log example6 &map[ctx_key:ctx_value]
//Jan 22 22:42:14 hades my_go_logger[113] <Emergency>: emergency Log example7 &map[ctx_key:ctx_value]
func ExampleSyslogHandler() {
	syslogHandler, _ := handler.NewSyslog(
		formatter.NewLine("%[2]s %[1]s %[3]s"),
		"",
		"",
		syslog.LOG_DEBUG,
		"my_go_logger")

	syslogHandler.Handle(logger.Entry{Message: "Log example", Level: logger.DebugLevel})
	syslogHandler.Handle(logger.Entry{Message: "Log example", Level: logger.InfoLevel})
	syslogHandler.Handle(logger.Entry{Message: "Log example", Level: logger.NoticeLevel})
	syslogHandler.Handle(logger.Entry{Message: "Log example", Level: logger.WarningLevel})
	syslogHandler.Handle(logger.Entry{Message: "Log example", Level: logger.ErrorLevel})
	syslogHandler.Handle(logger.Entry{Message: "Log example", Level: logger.CriticalLevel})
	syslogHandler.Handle(logger.Entry{Message: "Log example", Level: logger.AlertLevel})
	syslogHandler.Handle(logger.Entry{Message: "Log example", Level: logger.EmergencyLevel})
	//Output:
}

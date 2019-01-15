package example_handler_test

import (
	"os"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/formatter"
	"github.com/instabledesign/logger/handler"
)

func ExampleStreamHandler() {
	lineFormatter := formatter.NewLine("%[2]s | %[1]s")
	lineLogHandler := handler.NewStream(os.Stdout, lineFormatter)

	myLogger := logger.NewLogger(lineLogHandler)

	myLogger.Info("Log example", nil)

	//Output:
	// info | Log example
}

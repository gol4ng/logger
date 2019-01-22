package example_handler_test

import (
	"os"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/formatter"
	"github.com/instabledesign/logger/handler"
)

func ExampleStreamHandler() {
	lineFormatter := formatter.NewDefaultFormatter()
	lineLogHandler := handler.NewStream(os.Stdout, lineFormatter)
	lineLogHandler.Handle(logger.Entry{Message: "Log example"})

	//Output:
	// emergency Log example
}

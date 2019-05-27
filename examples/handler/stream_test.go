package example_handler_test

import (
	"os"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
)

func ExampleStreamHandler() {
	lineFormatter := formatter.NewDefaultFormatter()
	lineLogHandler := handler.Stream(os.Stdout, lineFormatter)
	lineLogHandler(logger.Entry{Message: "Log example"})

	//Output:
	//<emergency> Log example
}

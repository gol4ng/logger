package example_handler_test

import (
	"os"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
)

func ExampleGroupHandler() {
	lineFormatter := formatter.NewDefaultFormatter()
	lineLogHandler := handler.NewStream(os.Stdout, lineFormatter)

	jsonFormatter := formatter.NewJson()
	jsonLogHandler := handler.NewStream(os.Stdout, jsonFormatter)

	groupHandler := handler.NewGroup(lineLogHandler, jsonLogHandler)

	groupHandler.Handle(logger.Entry{Message: "Log example"})

	//Output:
	// emergency Log example
	// {"Message":"Log example","Level":0,"Context":null}
}

package example_handler_test

import (
	"os"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
)

func ExampleGroupHandler() {
	lineFormatter := formatter.NewDefaultFormatter()
	lineLogHandler := handler.Stream(os.Stdout, lineFormatter)

	jsonFormatter := formatter.NewJsonEncoder()
	jsonLogHandler := handler.Stream(os.Stdout, jsonFormatter)

	groupHandler := handler.Group(lineLogHandler, jsonLogHandler)

	groupHandler(logger.Entry{Message: "Log example"})

	//Output:
	// <emergency> Log example
	// {"Message":"Log example","Level":0,"Context":null}
}

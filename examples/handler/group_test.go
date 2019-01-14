// +build example

package example_handler_test

import (
	"os"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/formatter"
	"github.com/instabledesign/logger/handler"
)

func ExampleGroupHandler() {
	myString := "Log example"

	lineFormatter := formatter.NewLine("%[2]s | %[1]s")
	lineLogHandler := handler.NewStream(os.Stdout, lineFormatter)

	jsonFormatter := formatter.NewJson()
	jsonLogHandler := handler.NewStream(os.Stdout, jsonFormatter)

	groupHandler := handler.NewGroup([]logger.HandlerInterface{lineLogHandler, jsonLogHandler})

	myLogger := logger.Logger{HandlerInterface: groupHandler}

	myLogger.Info(myString, nil)

	//Output:
	// info | Log example
	// {"Message":"Log example","Level":0,"Context":null}
}

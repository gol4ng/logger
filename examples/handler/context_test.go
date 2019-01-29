package example_handler_test

import (
	"os"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
)

func ExampleContextHandler() {
	streamHandler := handler.NewStream(os.Stdout, formatter.NewDefaultFormatter())

	contextHandler := handler.NewContext(streamHandler, map[string]interface{}{"a": "1", "b": "2"})

	myLogger := logger.NewLogger(handler.NewGroup(contextHandler, streamHandler))

	myLogger.Debug("will be printed", &map[string]interface{}{"b": "3", "c": "4"})

	//Output:
	// debug will be printed
	// debug will be printed
}

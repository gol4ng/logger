package example_handler_test

import (
	"os"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/formatter"
	"github.com/instabledesign/logger/handler"
)

func ExampleContextHandler() {
	streamHandler := handler.NewStream(os.Stdout, formatter.NewLine("%s %s %s"))

	contextHandler := handler.NewContext(streamHandler, map[string]interface{}{"a": "1", "b": "2"})

	myLogger := logger.NewLogger(handler.NewGroup(contextHandler, streamHandler))

	myLogger.Debug("will be printed", &map[string]interface{}{"b": "3", "c": "4"})

	//Output:
	// will be printed debug &map[c:4 a:1 b:2]
	// will be printed debug &map[b:3 c:4]
}

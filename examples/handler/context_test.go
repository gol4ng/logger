package example_handler_test

import (
	"os"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
)

func ExampleContextHandler() {
	streamHandler := handler.NewStream(os.Stdout, formatter.NewJsonEncoder())

	contextHandler := handler.NewContext(streamHandler, logger.Ctx("my_value_1", "value 1"))

	myLogger := logger.NewLogger(handler.NewGroup(contextHandler, streamHandler))

	_ = myLogger.Debug("will be printed", logger.Ctx("my_value_1", "overwrited value 1"))

	_ = myLogger.Debug("only context handler values will be printed", nil)

	//Output:
	//{"Message":"will be printed","Level":7,"Context":{"my_value_1":"overwrited value 1"}}
	//{"Message":"will be printed","Level":7,"Context":{"my_value_1":"overwrited value 1"}}
	//{"Message":"only context handler values will be printed","Level":7,"Context":{"my_value_1":"value 1"}}
	//{"Message":"only context handler values will be printed","Level":7,"Context":null}
}

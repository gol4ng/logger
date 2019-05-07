package example_middleware_test

import (
	"github.com/gol4ng/logger/middleware"
	"os"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
)

func ExampleContextHandler() {
	streamHandler := handler.Stream(os.Stdout, formatter.NewJsonEncoder())

	contextHandler := middleware.Context(logger.Ctx("my_value_1", "value 1"))

	myLogger := logger.NewLogger(handler.Group(contextHandler(streamHandler), streamHandler))

	_ = myLogger.Debug("will be printed", logger.Ctx("my_value_1", "overwrited value 1"))

	_ = myLogger.Debug("only context handler values will be printed", nil)

	//Output:
	//{"Message":"will be printed","Level":7,"Context":{"my_value_1":"overwrited value 1"}}
	//{"Message":"will be printed","Level":7,"Context":{"my_value_1":"overwrited value 1"}}
	//{"Message":"only context handler values will be printed","Level":7,"Context":{"my_value_1":"value 1"}}
	//{"Message":"only context handler values will be printed","Level":7,"Context":null}
}

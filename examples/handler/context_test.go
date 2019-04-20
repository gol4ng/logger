package example_handler_test

import (
	"os"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
)

func ExampleContextHandler() {
	streamHandler := handler.NewStream(os.Stdout, formatter.NewJsonMarshall())

	//{"a": "1", "b": "2"}
	contextHandler := handler.NewContext(streamHandler, logger.NewContext().
		String("my_value_2", "context value 2").
		String("my_value_3", "context value 3"),
	)

	myLogger := logger.NewLogger(handler.NewGroup(contextHandler, streamHandler))

	_ = myLogger.Debug("will be printed", logger.NewContext().
		String("my_value_1", "value 1").
		String("my_value_2", "value 2"),
	)

	_ = myLogger.Debug("only context handler values will be printed", nil)

	//Output:
	// {"Message":"will be printed","Level":7,"Context":{"my_value_1":{"Type":16,"Value":"value 1"},"my_value_2":{"Type":16,"Value":"value 2"},"my_value_3":{"Type":16,"Value":"context value 3"}}}
	// {"Message":"will be printed","Level":7,"Context":{"my_value_1":{"Type":16,"Value":"value 1"},"my_value_2":{"Type":16,"Value":"value 2"}}}
	// {"Message":"only context handler values will be printed","Level":7,"Context":{"my_value_2":{"Type":16,"Value":"context value 2"},"my_value_3":{"Type":16,"Value":"context value 3"}}}
	// {"Message":"only context handler values will be printed","Level":7,"Context":null}
}

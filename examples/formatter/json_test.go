package example_formatter_test

import (
	"fmt"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
)

func ExampleJsonFormatter() {
	jsonFormatter := formatter.NewJSONEncoder()

	fmt.Println(jsonFormatter.Format(
		logger.Entry{
			Message: "My log message",
			Level: logger.InfoLevel,
			Context: logger.NewContext().Add("my_key", "my_value"),
		},
	))
	//Output:
	//{"Message":"My log message","Level":6,"Context":{"my_key":"my_value"}}
}

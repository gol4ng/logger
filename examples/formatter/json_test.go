package example_formatter_test

import (
	"fmt"

	"github.com/gol4ng/logger/formatter"

	"github.com/gol4ng/logger"
)

func ExampleJsonFormatter() {
	jsonFormatter := formatter.NewJsonMarshall()

	fmt.Println(jsonFormatter.Format(
		logger.Entry{
			Message: "My log message",
			Level: logger.InfoLevel,
			Context: logger.NewContext().String("my_key", "my_value"),
		},
	))
	//Output:
	// {"Message":"My log message","Level":6,"Context":{"my_key":{"Type":16,"Value":"my_value"}}}
}

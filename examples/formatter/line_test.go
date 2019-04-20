package example_formatter_test

import (
	"fmt"

	"github.com/gol4ng/logger/formatter"

	"github.com/gol4ng/logger"
)

func ExampleLineFormatter() {
	lineFormatter := formatter.NewLine("%s %s %s")

	//TODO fix serialization

	fmt.Println(lineFormatter.Format(
		logger.Entry{
			Message: "My log message",
			Level: logger.InfoLevel,
			Context: logger.NewContext().String("my_key", "my_value"),
		},
	))

	//Output:
	// My log message info &map[my_key:my_value]
}

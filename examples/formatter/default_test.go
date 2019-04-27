package example_formatter_test

import (
	"fmt"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
)

func ExampleDefaultFormatter() {
	defaultFormatter := formatter.NewDefaultFormatter()

	fmt.Println(defaultFormatter.Format(
		logger.Entry{
			Message: "My log message",
			Level: logger.InfoLevel,
			Context: logger.NewContext().Add("my_key", "my_value"),
		},
	))

	//Output:
	//<info> My log message {"my_key":"my_value"}
}

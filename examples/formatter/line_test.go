package example_formatter_test

import (
	"fmt"

	"github.com/gol4ng/logger/formatter"

	"github.com/gol4ng/logger"
)

func ExampleLineFormatter() {
	lineFormatter := formatter.NewLine("%s %s %s")

	fmt.Println(lineFormatter.Format(logger.Entry{Message: "My log message", Level: logger.InfoLevel, Context: &map[string]interface{}{"my_key": "my_value"}}))

	//Output:
	// My log message info &map[my_key:my_value]
}

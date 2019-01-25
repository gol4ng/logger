package example_formatter_test

import (
	"fmt"

	"github.com/gol4ng/logger/formatter"

	"github.com/gol4ng/logger"
)

func ExampleJsonFormatter() {
	jsonFormatter := formatter.NewJson()

	fmt.Println(jsonFormatter.Format(logger.Entry{Message: "My log message", Level: logger.InfoLevel, Context: &map[string]interface{}{"my_key": "my_value"}}))

	//Output:
	// {"Message":"My log message","Level":6,"Context":{"my_key":"my_value"}}
}

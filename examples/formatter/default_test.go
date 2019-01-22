package example_formatter_test

import (
	"fmt"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/formatter"
)

func ExampleDefaultFormatter() {
	jsonFormatter := formatter.NewDefaultFormatter()

	fmt.Println(jsonFormatter.Format(logger.Entry{Message: "My log message", Level: logger.InfoLevel, Context: &map[string]interface{}{"my_key": "my_value"}}))

	//Output:
	// info My log message
}

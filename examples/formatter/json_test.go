package example_formatter_test

import (
	"fmt"

	"github.com/instabledesign/logger/formatter"

	"github.com/instabledesign/logger"
)

func ExampleJsonFormatter() {
	jsonFormatter := formatter.NewJson()

	fmt.Println(jsonFormatter.Format(logger.Entry{Message: "My log message", Level: logger.InfoLevel, Context: &map[string]interface{}{"my_key": "my_value"}}))

	//Output:
	// {"Message":"My log message","Level":0,"Context":{"my_key":"my_value"}}
}

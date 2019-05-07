package example_middleware_test

import (
	"github.com/gol4ng/logger/middleware"
	"os"

	"github.com/gol4ng/logger/formatter"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
)

func ExampleMinLevelFilterHandler() {
	streamHandler := handler.Stream(os.Stdout, formatter.NewDefaultFormatter())

	minLvlFilterHandler := middleware.MinLevelFilter(logger.WarningLevel)(streamHandler)
	minLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.DebugLevel})
	minLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.InfoLevel})
	minLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.NoticeLevel})
	minLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.WarningLevel})
	minLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.ErrorLevel})
	minLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.CriticalLevel})
	minLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.AlertLevel})
	minLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.EmergencyLevel})

	//Output:
	//<warning> will be printed
	//<error> will be printed
	//<critical> will be printed
	//<alert> will be printed
	//<emergency> will be printed
}

func ExampleRangeLevelFilterHandler() {
	streamHandler := handler.Stream(os.Stdout, formatter.NewDefaultFormatter())

	rangeLvlFilterHandler := middleware.RangeLevelFilter(logger.InfoLevel, logger.WarningLevel)(streamHandler)

	rangeLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.DebugLevel})
	rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.InfoLevel})
	rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.NoticeLevel})
	rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.WarningLevel})
	rangeLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.ErrorLevel})
	rangeLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.CriticalLevel})
	rangeLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.AlertLevel})
	rangeLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.EmergencyLevel})

	//Output:
	//<info> will be printed
	//<notice> will be printed
	//<warning> will be printed
}

func ExampleCustomFilterHandler() {
	streamHandler := handler.Stream(os.Stdout, formatter.NewDefaultFormatter())

	rangeLvlFilterHandler := middleware.Filter(func(e logger.Entry) bool {
		return e.Level == logger.InfoLevel || e.Level == logger.AlertLevel
	})(streamHandler)

	rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.DebugLevel})
	rangeLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.InfoLevel})
	rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.NoticeLevel})
	rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.WarningLevel})
	rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.ErrorLevel})
	rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.CriticalLevel})
	rangeLvlFilterHandler(logger.Entry{Message: "will be excluded", Level: logger.AlertLevel})
	rangeLvlFilterHandler(logger.Entry{Message: "will be printed", Level: logger.EmergencyLevel})

	//Output:
	//<debug> will be printed
	//<notice> will be printed
	//<warning> will be printed
	//<error> will be printed
	//<critical> will be printed
	//<emergency> will be printed
}

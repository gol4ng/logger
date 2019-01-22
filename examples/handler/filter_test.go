package example_handler_test

import (
	"os"

	"github.com/instabledesign/logger/formatter"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/handler"
)

func ExampleMinLevelFilterHandler() {
	streamHandler := handler.NewStream(os.Stdout, formatter.NewLine("%[2]s"))

	minLvlFilterHandler := handler.NewMinLevelFilter(streamHandler, logger.WarningLevel)
	minLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.DebugLevel})
	minLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.InfoLevel})
	minLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.NoticeLevel})
	minLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.WarningLevel})
	minLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.ErrorLevel})
	minLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.CriticalLevel})
	minLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.AlertLevel})
	minLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.EmergencyLevel})

	//Output:
	//warning
	//error
	//critical
	//alert
	//emergency
}

func ExampleRangeLevelFilterHandler() {
	streamHandler := handler.NewStream(os.Stdout, formatter.NewLine("%[2]s"))

	rangeLvlFilterHandler := handler.NewRangeLevelFilter(streamHandler, logger.InfoLevel, logger.WarningLevel)

	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.DebugLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.InfoLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.NoticeLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.WarningLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.ErrorLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.CriticalLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.AlertLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.EmergencyLevel})

	//Output:
	//info
	//notice
	//warning
}

func ExampleCustomFilterHandler() {
	streamHandler := handler.NewStream(os.Stdout, formatter.NewLine("%[2]s"))

	rangeLvlFilterHandler := handler.NewFilter(streamHandler, func(e logger.Entry) bool {
		return e.Level == logger.InfoLevel || e.Level == logger.AlertLevel
	})

	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.DebugLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.InfoLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.NoticeLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.WarningLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.ErrorLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.CriticalLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.AlertLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.EmergencyLevel})

	//Output:
	//debug
	//notice
	//warning
	//error
	//critical
	//emergency
}

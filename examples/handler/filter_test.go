package example_handler_test

import (
	"os"

	"github.com/gol4ng/logger/formatter"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
)

func ExampleMinLevelFilterHandler() {
	streamHandler := handler.NewStream(os.Stdout, formatter.NewDefaultFormatter())

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
	//<warning> will be printed
	//<error> will be printed
	//<critical> will be printed
	//<alert> will be printed
	//<emergency> will be printed
}

func ExampleRangeLevelFilterHandler() {
	streamHandler := handler.NewStream(os.Stdout, formatter.NewDefaultFormatter())

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
	//<info> will be printed
	//<notice> will be printed
	//<warning> will be printed
}

func ExampleCustomFilterHandler() {
	streamHandler := handler.NewStream(os.Stdout, formatter.NewDefaultFormatter())

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
	//<debug> will be printed
	//<notice> will be printed
	//<warning> will be printed
	//<error> will be printed
	//<critical> will be printed
	//<emergency> will be printed
}

package example_handler_test

import (
	"os"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/handler"
)

func ExampleMinLevelFilterHandler() {
	streamHandler := handler.NewStream(os.Stdout, logger.NewNilFormatter())

	minLvlFilterHandler := handler.NewMinLevelFilter(streamHandler, logger.WarnLevel)
	minLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.DebugLevel})
	minLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.InfoLevel})
	minLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.WarnLevel})
	minLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.ErrorLevel})
	minLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.PanicLevel})
	minLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.FatalLevel})

	//Output:
	// {will be printed warn <nil>}
	// {will be printed error <nil>}
	// {will be printed panic <nil>}
	// {will be printed fatal <nil>}
}

func ExampleRangeLevelFilterHandler() {
	streamHandler := handler.NewStream(os.Stdout, logger.NewNilFormatter())

	rangeLvlFilterHandler := handler.NewRangeLevelFilter(streamHandler, logger.InfoLevel, logger.WarnLevel)

	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.DebugLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.InfoLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.WarnLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.ErrorLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.PanicLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.FatalLevel})

	//Output:
	// {will be printed info <nil>}
	// {will be printed warn <nil>}
}

func ExampleCustomFilterHandler() {
	streamHandler := handler.NewStream(os.Stdout, logger.NewNilFormatter())

	rangeLvlFilterHandler := handler.NewFilter(streamHandler, func(e logger.Entry) bool {
		return e.Level == logger.InfoLevel || e.Level == logger.PanicLevel
	})

	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.DebugLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.InfoLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.WarnLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.ErrorLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be excluded", Level: logger.PanicLevel})
	rangeLvlFilterHandler.Handle(logger.Entry{Message: "will be printed", Level: logger.FatalLevel})

	//Output:
	// {will be printed debug <nil>}
	// {will be printed warn <nil>}
	// {will be printed error <nil>}
	// {will be printed fatal <nil>}
}

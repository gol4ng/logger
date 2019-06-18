package logger

// logger.HandlerInterface allows you to process a logger entry
// (i.e: transform message to GELF, send the message to slack ...)
// You can find basic handler implementations in the handler package.
type HandlerInterface func(Entry) error

var NopHandler HandlerInterface = func(entry Entry) error {
	return nil
}

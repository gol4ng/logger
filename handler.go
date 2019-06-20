package logger

// HandlerInterface allows you to process a Entry
// See basic handler implementations in the handler package.
type HandlerInterface func(Entry) error

// NopHandler is a no operating handler that do nothing with received Entry
var NopHandler HandlerInterface = func(entry Entry) error {
	return nil
}

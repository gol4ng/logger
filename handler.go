package logger

type HandlerInterface func(Entry) error

var NopHandler HandlerInterface = func(entry Entry) error {
	return nil
}

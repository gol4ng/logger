package logger

type HandlerInterface interface {
	Handle(e Entry) error
}

type NopHandler struct{}

func (n *NopHandler) Handle(entry Entry) error {
	return nil
}

func NewNopHandler() *NopHandler {
	return &NopHandler{}
}

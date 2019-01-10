package logger

type HandlerInterface interface {
	Handle(e Entry) error
}

type NilHandler struct{}

func (n *NilHandler) Handle(e Entry) error {
	return nil
}

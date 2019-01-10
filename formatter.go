package logger

type FormatterInterface interface {
	Format(entry Entry) interface{}
}

type NilFormatter struct{}

func (n *NilFormatter) Format(e Entry) interface{} {
	return e
}

func NewNilFormatter() *NilFormatter {
	return &NilFormatter{}
}

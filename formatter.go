package logger

type FormatterInterface interface {
	Format(entry Entry) string
}

type NilFormatter struct{}

func (n *NilFormatter) Format(e Entry) string {
	return "[" + string(e.Level) + "]" + e.Message
}

func NewNilFormatter() *NilFormatter {
	return &NilFormatter{}
}

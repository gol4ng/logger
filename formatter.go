package logger

type FormatterInterface interface {
	Format(entry Entry) string
}

type NopFormatter struct{}

func (n *NopFormatter) Format(entry Entry) string {
	return ""
}

func NewNopFormatter() *NopFormatter {
	return &NopFormatter{}
}

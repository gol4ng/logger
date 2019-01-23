package logger

type FormatterInterface interface {
	Format(entry Entry) string
}

type NopFormatter struct{}

func (n *NopFormatter) Format(e Entry) string {
	return ""
}

func NewNopFormatter() *NopFormatter {
	return &NopFormatter{}
}

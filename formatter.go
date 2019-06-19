package logger

// allows you to format a logger entry
// ex: format a log entry to gelf format
type FormatterInterface interface {
	Format(entry Entry) string
}

type NopFormatter struct{}

func (n *NopFormatter) Format(entry Entry) string {
	return ""
}

// instantiate a formatter that returns a void string
func NewNopFormatter() *NopFormatter {
	return &NopFormatter{}
}

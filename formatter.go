package logger

// FormatterInterface will convert Entry into string
// ex: format a log entry to string/json/... format
type FormatterInterface interface {
	Format(entry Entry) string
}

// NopFormatter is a no operating formatter
type NopFormatter struct{}

// Format will return empty string
func (n *NopFormatter) Format(_ Entry) string {
	return ""
}

// NewNopFormatter will create a NopFormatter
func NewNopFormatter() *NopFormatter {
	return &NopFormatter{}
}

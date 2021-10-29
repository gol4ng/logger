package formatter

import (
	"github.com/gol4ng/logger"
	"github.com/valyala/bytebufferpool"
)

var falseCondition = func(entry logger.Entry) bool {
	return false
}

// DefaultFormatter is the default Entry formatter
type DefaultFormatter struct {
	colored        func(entry logger.Entry) bool
	displayContext func(entry logger.Entry) bool
}

// Format will return Entry as string
func (n *DefaultFormatter) Format(entry logger.Entry) string {
	byteBuffer := bytebufferpool.Get()
	defer bytebufferpool.Put(byteBuffer)

	colored := n.colored(entry)
	if colored {
		switch entry.Level {
		case logger.DebugLevel:
			byteBuffer.WriteString("\x1b[1;36m")
		case logger.InfoLevel:
			byteBuffer.WriteString("\x1b[1;32m")
		case logger.NoticeLevel:
			byteBuffer.WriteString("\x1b[1;34m")
		case logger.WarningLevel:
			byteBuffer.WriteString("\x1b[1;33m")
		case logger.ErrorLevel:
			byteBuffer.WriteString("\x1b[1;31m")
		case logger.CriticalLevel:
			byteBuffer.WriteString("\x1b[1;30;47m")
		case logger.AlertLevel:
			byteBuffer.WriteString("\x1b[1;30;43m")
		case logger.EmergencyLevel:
			byteBuffer.WriteString("\x1b[1;37;41m")
		}
	}

	byteBuffer.WriteString(`<`)
	byteBuffer.WriteString(entry.Level.String())
	byteBuffer.WriteString(`>`)
	if colored {
		byteBuffer.WriteString("\x1b[m")
	}
	if entry.Message != "" {
		byteBuffer.WriteString(` `)
		byteBuffer.WriteString(entry.Message)
	}
	if entry.Context != nil && n.displayContext(entry) {
		byteBuffer.WriteString(` `)
		ContextToJSON(entry.Context, byteBuffer)
	}
	return byteBuffer.String()
}

// NewDefaultFormatter will create a new DefaultFormatter
func NewDefaultFormatter(options ...Option) *DefaultFormatter {
	f := &DefaultFormatter{
		colored: falseCondition,
		displayContext: falseCondition,
	}
	for _, option := range options {
		option(f)
	}
	return f
}

// Option is the option pattern interface for the DefaultFormatter
type Option func(*DefaultFormatter)

// WithColor function will enable ANSI colored formatting
func WithColor(enable bool) Option {
	return WithConditionalColor(func(_ logger.Entry) bool {
		return enable
	})
}

// WithConditionalColor function will enable ANSI colored formatting
func WithConditionalColor(conditional func(_ logger.Entry) bool) Option {
	return func(formatter *DefaultFormatter) {
		formatter.colored = conditional
	}
}

// WithContext function will display context printing
func WithContext(enable bool) Option {
	return WithConditionalContext(func(_ logger.Entry) bool {
		return enable
	})
}

// WithConditionalContext function will display context printing
func WithConditionalContext(conditional func(_ logger.Entry) bool) Option {
	return func(formatter *DefaultFormatter) {
		formatter.displayContext = conditional
	}
}

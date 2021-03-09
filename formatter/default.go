package formatter

import (
	"strings"

	"github.com/gol4ng/logger"
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
	n.init()
	builder := &strings.Builder{}

	colored := n.colored(entry)
	if colored {
		switch entry.Level {
		case logger.DebugLevel:
			builder.WriteString("\x1b[1;36m")
		case logger.InfoLevel:
			builder.WriteString("\x1b[1;32m")
		case logger.NoticeLevel:
			builder.WriteString("\x1b[1;34m")
		case logger.WarningLevel:
			builder.WriteString("\x1b[1;33m")
		case logger.ErrorLevel:
			builder.WriteString("\x1b[1;31m")
		case logger.CriticalLevel:
			builder.WriteString("\x1b[1;30;47m")
		case logger.AlertLevel:
			builder.WriteString("\x1b[1;30;43m")
		case logger.EmergencyLevel:
			builder.WriteString("\x1b[1;37;41m")
		}
	}

	builder.WriteString("<")
	builder.WriteString(entry.Level.String())
	builder.WriteString(">")
	if colored {
		builder.WriteString("\x1b[m")
	}
	if entry.Message != "" {
		builder.WriteString(" ")
		builder.WriteString(entry.Message)
	}
	if entry.Context != nil && n.displayContext(entry) {
		builder.WriteString(" ")
		ContextToJSON(entry.Context, builder)
	}
	return builder.String()
}

func (n *DefaultFormatter) init() {
	if n.colored == nil {
		n.colored = falseCondition
	}
	if n.displayContext == nil {
		n.displayContext = falseCondition
	}
}

// NewDefaultFormatter will create a new DefaultFormatter
func NewDefaultFormatter(options ...Option) *DefaultFormatter {
	f := &DefaultFormatter{}
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

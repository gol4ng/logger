package logger

import "context"

type contextKey int

const (
	loggerKey contextKey = iota
)

// InjectInContext will inject a logger into the go-context
func InjectInContext(ctx context.Context, l LoggerInterface) context.Context {
	return context.WithValue(ctx, loggerKey, l)
}

// FromContext will retrieve a logger from the go-context or return defaultLogger
func FromContext(ctx context.Context, defaultLogger LoggerInterface) LoggerInterface {
	if _logger, ok := ctx.Value(loggerKey).(LoggerInterface); ok {
		return _logger
	}
	return defaultLogger
}

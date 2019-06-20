package logger

import "context"

var contextKey struct{}

// InjectInContext will inject a logger into the go-context
func InjectInContext(ctx context.Context, l LoggerInterface) context.Context {
	return context.WithValue(ctx, contextKey, l)
}

// FromContext will retrieve a logger from the go-context or return defaultLogger
func FromContext(ctx context.Context, defaultLogger LoggerInterface) LoggerInterface {
	if _logger, ok := ctx.Value(contextKey).(LoggerInterface); ok {
		return _logger
	}
	return defaultLogger
}

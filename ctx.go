package logger

import "context"

var ContextKey struct{}

func InjectInContext(ctx context.Context, l LoggerInterface) context.Context {
	return context.WithValue(ctx, ContextKey, l)
}

func FromContext(ctx context.Context, l LoggerInterface) LoggerInterface {
	if _logger, ok := ctx.Value(ContextKey).(LoggerInterface); ok {
		return _logger
	}
	return l
}

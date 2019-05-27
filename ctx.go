package logger

import "context"

var contextKey struct{}

func InjectInContext(ctx context.Context, l LoggerInterface) context.Context {
	return context.WithValue(ctx, contextKey, l)
}

func FromContext(ctx context.Context, l LoggerInterface) LoggerInterface {
	if _logger, ok := ctx.Value(contextKey).(LoggerInterface); ok {
		return _logger
	}
	return l
}

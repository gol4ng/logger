package logger

import "context"

var contextKey struct{}
// inject a logger into the go-context
func InjectInContext(ctx context.Context, l LoggerInterface) context.Context {
	return context.WithValue(ctx, contextKey, l)
}
// retrieve a logger from the go-context if any
func FromContext(ctx context.Context, l LoggerInterface) LoggerInterface {
	if _logger, ok := ctx.Value(contextKey).(LoggerInterface); ok {
		return _logger
	}
	return l
}

package logger

import "context"

type ctxKey struct{}

func (l *Logger) WithContext(ctx context.Context) context.Context {
	if ctxL, ok := ctx.Value(ctxKey{}).(*Logger); ok {
		if ctxL == l {
			return ctx
		}
	}

	return context.WithValue(ctx, ctxKey{}, l)
}

func Ctx(ctx context.Context) *Logger {
	if l, ok := ctx.Value(ctxKey{}).(*Logger); ok {
		return l
	}
	return NewNopLogger()
}

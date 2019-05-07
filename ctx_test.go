package logger_test

import (
	"context"
	"github.com/gol4ng/logger"
	"gotest.tools/assert"
	"testing"
)

func Test_FromContext_Empty(t *testing.T) {
	log := logger.NewNopLogger()

	logCtx := logger.FromContext(context.Background(), log)

	assert.Equal(t, log, logCtx)
}

func Test_FromContext_AlreadyHave(t *testing.T) {
	ctxLog := logger.NewNopLogger()
	ctxWithLogger := logger.InjectInContext(context.Background(), ctxLog)

	log := logger.NewNopLogger()
	log2 := logger.FromContext(ctxWithLogger, log)
	assert.Equal(t, log2, ctxLog)
}

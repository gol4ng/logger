package logger_test

import (
	"context"
	"testing"

	"gotest.tools/assert"

	"github.com/gol4ng/logger"
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

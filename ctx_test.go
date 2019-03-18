package logger_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/gol4ng/logger"
)

func TestCtx(t *testing.T) {
	log := logger.NewLogger(logger.NewNopHandler())
	ctx := log.WithContext(context.Background())
	log2 := logger.Ctx(ctx)
	if !reflect.DeepEqual(log, log2) {
		t.Error("Ctx did not return the expected logger")
	}

	ctx = log.WithContext(ctx)
	log2 = logger.Ctx(ctx)
	if !reflect.DeepEqual(log, log2) {
		t.Error("Ctx did not return the expected logger")
	}

	log2 = logger.Ctx(context.Background())
	if !reflect.DeepEqual(logger.NewNopLogger(), log2) {
		t.Error("Ctx did not return the expected logger")
	}
}

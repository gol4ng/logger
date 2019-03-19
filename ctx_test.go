package logger_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/gol4ng/logger"
)

func Test_WithContext_New(t *testing.T) {
	log := logger.NewLogger(logger.NewNopHandler())
	ctx := log.WithContext(context.Background())
	log2 := logger.Ctx(ctx)
	if !reflect.DeepEqual(log, log2) {
		t.Error("Ctx did not return the expected logger")
	}
}

func Test_WithContext_AlreadyHave(t *testing.T) {
	log := logger.NewLogger(logger.NewNopHandler())
	ctx := log.WithContext(log.WithContext(context.Background()))
	log2 := logger.Ctx(ctx)
	if !reflect.DeepEqual(log, log2) {
		t.Error("Ctx did not return the expected logger")
	}
}

func Test_Ctx_Empty(t *testing.T) {
	log := logger.Ctx(context.Background())
	if !reflect.DeepEqual(logger.NewNopLogger(), log) {
		t.Error("Ctx did not return the expected logger")
	}
}

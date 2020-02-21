package testing_test

import (
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/middleware"
	"github.com/stretchr/testify/assert"

	testing_logger "github.com/gol4ng/logger/testing"
)

func TestNewLogger(t *testing.T) {
	myLogger, store := testing_logger.NewLogger()

	myLogger.Debug("test Debug")
	myLogger.Info("test Info", logger.Any("my_ctx_value", "value"))

	entries := store.GetEntries()
	assert.Len(t, entries, 2)

	entry1 := entries[0]
	assert.Equal(t, logger.Context{}, *entry1.Context)
	assert.Equal(t, logger.DebugLevel, entry1.Level)
	assert.Equal(t, "test Debug", entry1.Message)

	entry2 := entries[1]
	eCtx2 := *entry2.Context
	assert.Equal(t, "value", eCtx2["my_ctx_value"].Value)
	assert.Equal(t, logger.InfoLevel, entry2.Level)
	assert.Equal(t, "test Info", entry2.Message)
}

func TestNewLoggerRace(t *testing.T) {
	myLogger, store := testing_logger.NewLogger()

	go func() {
		myLogger.Debug("test Debug")
	}()

	go func() {
		store.GetEntries()
	}()
}

func TestNewLogger_GetEntries(t *testing.T) {
	myLogger, store := testing_logger.NewLogger()

	myLogger.Debug("test Debug")
	myLogger.Info("test Info", logger.Any("my_ctx_value", "value"))

	entries := store.GetEntries()
	assert.Len(t, entries, 2)
}

func TestNewLogger_CleanEntries(t *testing.T) {
	myLogger, store := testing_logger.NewLogger()

	myLogger.Debug("test Debug")
	myLogger.Info("test Info", logger.Any("my_ctx_value", "value"))

	entries := store.GetEntries()
	assert.Len(t, entries, 2)

	store.CleanEntries()

	assert.Empty(t, store.GetEntries())
}

func TestNewLogger_GetAndCleanEntries(t *testing.T) {
	myLogger, store := testing_logger.NewLogger()

	myLogger.Debug("test Debug")
	myLogger.Info("test Info", logger.Any("my_ctx_value", "value"))

	entries := store.GetAndCleanEntries()
	assert.Len(t, entries, 2)

	assert.Empty(t, store.GetEntries())
}

func TestNewLogger_Wrap(t *testing.T) {
	myLogger, store := testing_logger.NewLogger()

	myLogger.Wrap(middleware.Context(logger.Ctx("my_ctx_value", "value")))
	myLogger.Debug("test Debug")
	myLogger.Info("test Info")

	entries := store.GetEntries()
	assert.Len(t, entries, 2)

	for _, e := range entries {
		assert.Equal(t, "value", (*e.Context)["my_ctx_value"].Value)
	}

	entry1 := entries[0]
	assert.Equal(t, logger.DebugLevel, entry1.Level)
	assert.Equal(t, "test Debug", entry1.Message)

	entry2 := entries[1]
	assert.Equal(t, logger.InfoLevel, entry2.Level)
	assert.Equal(t, "test Info", entry2.Message)
}

func TestNewLogger_WrapNew(t *testing.T) {
	myLogger, store := testing_logger.NewLogger()

	defaultContext := logger.Context(map[string]logger.Field{
		"my_key": {Value: "my_value"},
	})
	myLogger2 := myLogger.WrapNew(middleware.Context(&defaultContext))

	myLogger.Debug("test Debug")
	myLogger.Info("test Info")

	myLogger2.Debug("test wrapped Debug")
	myLogger2.Info("test wrapped Info")

	entries := store.GetEntries()
	assert.Len(t, entries, 4)

	entry1 := entries[0]
	assert.Equal(t, logger.Context{}, *entry1.Context)
	assert.Equal(t, logger.DebugLevel, entry1.Level)
	assert.Equal(t, "test Debug", entry1.Message)

	entry2 := entries[1]
	assert.Equal(t, logger.Context{}, *entry2.Context)
	assert.Equal(t, logger.InfoLevel, entry2.Level)
	assert.Equal(t, "test Info", entry2.Message)

	entry3 := entries[2]
	entry3Ctx := *entry3.Context
	assert.Equal(t, "my_value", entry3Ctx["my_key"].Value)
	assert.Equal(t, logger.DebugLevel, entry3.Level)
	assert.Equal(t, "test wrapped Debug", entry3.Message)

	entry4 := entries[3]
	entry4Ctx := *entry4.Context
	assert.Equal(t, "my_value", entry4Ctx["my_key"].Value)
	assert.Equal(t, logger.InfoLevel, entry4.Level)
	assert.Equal(t, "test wrapped Info", entry4.Message)
}

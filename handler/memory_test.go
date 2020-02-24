package handler_test

import (
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
	"github.com/stretchr/testify/assert"
)

func TestMemory_Handle(t *testing.T) {
	entry := logger.Entry{}

	h := &handler.Memory{}

	err := h.Handle(entry)
	assert.NoError(t, err)

	entries := h.GetEntries()
	assert.Len(t, entries, 1)

	assert.Equal(t, entry, entries[0])
}

func TestMemory_HandleRace(t *testing.T) {
	h := &handler.Memory{}

	go func() {
		err := h.Handle(logger.Entry{})
		assert.NoError(t, err)
	}()

	go func() {
		h.GetEntries()
	}()
}

func TestNewLogger_CleanEntries(t *testing.T) {
	entry := logger.Entry{}

	h := &handler.Memory{}

	err := h.Handle(entry)
	assert.NoError(t, err)

	entries := h.GetEntries()
	assert.Len(t, entries, 1)

	assert.Equal(t, entry, entries[0])

	h.CleanEntries()

	assert.Empty(t, h.GetEntries())
}

func TestNewLogger_GetAndCleanEntries(t *testing.T) {
	entry := logger.Entry{}

	h := &handler.Memory{}

	err := h.Handle(entry)
	assert.NoError(t, err)

	entries := h.GetAndCleanEntries()
	assert.Len(t, entries, 1)

	assert.Equal(t, entry, entries[0])

	assert.Empty(t, h.GetEntries())
}

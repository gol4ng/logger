package provider_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger/writer"
	"github.com/gol4ng/logger/writer/provider"
)

func TestCompressProvider_WithError(t *testing.T) {
	var writer1 io.Writer
	var MyProvider = func(w io.Writer) (io.Writer, error) {
		assert.Equal(t, writer1, w)
		return w, errors.New("my_fake_provider_error")
	}
	w := provider.CompressProvider(MyProvider)
	file, err := w(writer1)
	assert.Nil(t, file)
	assert.EqualError(t, err, "my_fake_provider_error")
}

func TestCompressProvider(t *testing.T) {
	createdFile := &os.File{}
	var MyProvider = func(w io.Writer) (io.Writer, error) {
		assert.Nil(t, w)
		return createdFile, nil
	}

	w := provider.CompressProvider(MyProvider)
	newFile, err := w(nil)
	assert.Nil(t, err)
	assert.IsType(t, &writer.CompressWriter{}, newFile)
}

package provider

import (
	"io"

	"github.com/gol4ng/logger/writer"
)

// CompressProvider will decorate writer.Provider with compression
func CompressProvider(provider writer.Provider, options ...writer.CompressOption) writer.Provider {
	return func(w io.Writer) (io.Writer, error) {
		file, err := provider(w)
		if err != nil {
			return nil, err
		}

		return writer.NewCompressWriter(file, options...), nil
	}
}

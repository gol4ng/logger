package provider

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gol4ng/logger/writer"
)

func LogFileProvider(name string, format string, timeFormat string) writer.Provider {
	basePath := fmt.Sprintf(format, name)
	return func(w io.Writer) (io.Writer, error) {
		if w != nil {
			if err := os.Rename(basePath, fmt.Sprintf(format, time.Now().Format(timeFormat))); err != nil {
				return nil, err
			}
			if _w, ok := w.(io.Closer); ok {
				if err := _w.Close(); err != nil {
					return nil, err
				}
			}
		}
		return os.OpenFile(
			basePath,
			os.O_CREATE|os.O_APPEND|os.O_WRONLY,
			0666,
		)
	}
}

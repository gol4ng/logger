package provider

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gol4ng/logger/writer"
)

func TimeFileProvider(format string, timeFormat string) writer.Provider {
	return func(w io.Writer) (io.Writer, error) {
		if w != nil {
			if _w, ok := w.(io.Closer); ok {
				if err := _w.Close(); err != nil {
					return nil, err
				}
			}
		}

		return os.OpenFile(
			fmt.Sprintf(format, time.Now().Format(timeFormat)),
			os.O_CREATE|os.O_APPEND|os.O_WRONLY,
			0666,
		)
	}
}

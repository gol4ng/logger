package provider

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gol4ng/logger/writer"
)

// TimeFileProvider provide a new file (io.Writer) having named with a given filename format
// ex: `provider.TimeFileProvider("fake_format_%s", "2006-01-02")` will create a file named "fake_format_2019-08-17" (i.e fake_format_<ISO_8601_DATE>)
func TimeFileProvider(fileNameFormat string, timeFormat string) writer.Provider {
	return func(w io.Writer) (io.Writer, error) {
		// if the provider already have an open io.writer when it is called, close it
		// (i.e: close current file and create a new one)
		if w != nil {
			if _w, ok := w.(io.Closer); ok {
				if err := _w.Close(); err != nil {
					return nil, err
				}
			}
		}
		// create a file (io.Writer) having the given filename format
		return os.OpenFile(
			fmt.Sprintf(fileNameFormat, time.Now().Format(timeFormat)),
			os.O_CREATE|os.O_APPEND|os.O_WRONLY,
			0666,
		)
	}
}

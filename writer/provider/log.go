package provider

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gol4ng/logger/writer"
)

// LogFileProvider creates one static file with the given name and format
// if the file name already exists, it renames it with the time and create a new static file
//
// ex: LogFileProvider("dev", os.TempDir()+"/my_app_%s.log", "2006-01-02") will create a file named "/tmp/my_app_dev.log"
// and each time the provider is called, it will backup the current "/tmp/dev.log" to "/tmp/my_app_2006-01-02.log"
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

package file_provider

import (
	"fmt"
	"os"
	"time"

	"github.com/gol4ng/logger/writer"
)

func LogFileProvider(name string, format string, timeFormat string) writer.FileProvider {
	basePath := fmt.Sprintf(format, name)
	return func(f *os.File) (*os.File, error) {
		if f != nil {
			err := os.Rename(basePath, fmt.Sprintf(format, time.Now().Format(timeFormat)))
			if err != nil {
				return nil, err
			}
			err = f.Close()
			if err != nil {
				return nil, err
			}
		}
		return os.OpenFile(
			basePath,
			os.O_CREATE|os.O_APPEND|os.O_WRONLY,
			0666,
		)
	}
}

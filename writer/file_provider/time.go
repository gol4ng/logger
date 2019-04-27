package file_provider

import (
	"fmt"
	"os"
	"time"

	"github.com/gol4ng/logger/writer"
)

func TimeFileProvider(format string, timeFormat string) writer.FileProvider {
	return func(f *os.File) (*os.File, error) {
		if f != nil {
			err := f.Close()
			if err != nil {
				return nil, err
			}
		}
		return os.OpenFile(
			fmt.Sprintf(format, time.Now().Format(timeFormat)),
			os.O_CREATE|os.O_APPEND|os.O_WRONLY,
			0666,
		)
	}
}

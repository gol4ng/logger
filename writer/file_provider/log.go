package file_provider

import (
	"fmt"
	"os"
	"time"
)

func LogFileProvider(name string, format string, timeFormat string) func(f *os.File) (*os.File, error) {
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
		return os.Create(basePath)
	}
}

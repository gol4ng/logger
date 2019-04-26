package writer

import "os"

type FileProvider func(*os.File) (*os.File, error)

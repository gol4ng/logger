package writer

import (
	"io"
)

type Provider func(io.Writer) (io.Writer, error)

package writer

import (
	"io"
)

// Provider is an io.Writer provider it is used to get an io.Writer
// Implementation should manage where and how kind of Writer it will return
// See provider implementations in the provider package.
type Provider func(io.Writer) (io.Writer, error)

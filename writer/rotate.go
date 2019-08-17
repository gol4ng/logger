package writer

import (
	"io"
)

// RotateWriter is a Writer with extra rotate feature that allows you to change the writer as you need
// See rotate writer implementations in the writer package.
type RotateWriter interface {
	io.Writer
	Rotate() error
}

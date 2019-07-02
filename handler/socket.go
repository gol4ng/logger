package handler

import (
	"bytes"
	"net"

	"github.com/gol4ng/logger"
)

// Socket will create a an handler that format and send data into a connection
func Socket(connection net.Conn, formatter logger.FormatterInterface) logger.HandlerInterface {
	return func(entry logger.Entry) error {
		buffer := bytes.NewBuffer([]byte(formatter.Format(entry)))
		_, err := connection.Write(buffer.Bytes())
		return err
	}
}

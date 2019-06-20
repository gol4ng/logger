package handler

import (
	"log/syslog"

	"github.com/gol4ng/logger"
)

// Syslog will format and send Entry to a syslog server
func Syslog(formatter logger.FormatterInterface, network, raddr string, priority syslog.Priority, tag string) (logger.HandlerInterface, error) {
	if "" == tag {
		tag = "golang"
	}
	writer, err := syslog.Dial(network, raddr, priority, tag)
	if err != nil {
		return nil, err
	}

	return func(entry logger.Entry) error {
		msg := formatter.Format(entry)
		switch entry.Level {
		case logger.DebugLevel:
			return writer.Debug(msg)
		case logger.InfoLevel:
			return writer.Info(msg)
		case logger.NoticeLevel:
			return writer.Notice(msg)
		case logger.WarningLevel:
			return writer.Warning(msg)
		case logger.ErrorLevel:
			return writer.Err(msg)
		case logger.CriticalLevel:
			return writer.Crit(msg)
		case logger.AlertLevel:
			return writer.Alert(msg)
		case logger.EmergencyLevel:
			return writer.Emerg(msg)
		default:
			_, err := writer.Write([]byte(msg))
			return err
		}
	}, nil
}

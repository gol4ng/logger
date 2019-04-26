package handler

import (
	"log/syslog"

	"github.com/gol4ng/logger"
)

type Syslog struct {
	writer    *syslog.Writer
	formatter logger.FormatterInterface
}

func (s *Syslog) Handle(entry logger.Entry) error {
	msg := s.formatter.Format(entry)
	switch entry.Level {
	case logger.DebugLevel:
		return s.writer.Debug(msg)
	case logger.InfoLevel:
		return s.writer.Info(msg)
	case logger.NoticeLevel:
		return s.writer.Notice(msg)
	case logger.WarningLevel:
		return s.writer.Warning(msg)
	case logger.ErrorLevel:
		return s.writer.Err(msg)
	case logger.CriticalLevel:
		return s.writer.Crit(msg)
	case logger.AlertLevel:
		return s.writer.Alert(msg)
	case logger.EmergencyLevel:
		return s.writer.Emerg(msg)
	default:
		_, err := s.writer.Write([]byte(msg))
		return err
	}
}

func NewSyslog(formatter logger.FormatterInterface, network, raddr string, priority syslog.Priority, tag string) (*Syslog, error) {
	if "" == tag {
		tag = "golang"
	}
	w, err := syslog.Dial(network, raddr, priority, tag)

	return &Syslog{writer: w, formatter: formatter}, err
}

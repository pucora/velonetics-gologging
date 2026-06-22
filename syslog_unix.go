//go:build !windows

package gologging

import (
	"io"
	"log/syslog"
	"strings"
)

const (
	defaultSyslogFacility = syslog.LOG_LOCAL3
	defaultSyslogSeverity = syslog.LOG_CRIT
)

type syslogPriority = syslog.Priority

func newSyslogWriter(severity, facility syslog.Priority, prefix string) (io.Writer, error) {
	w, err := syslog.New(severity|facility, prefix)
	if err != nil {
		return nil, err
	}
	return syslogIoWriterWrapper{w}, nil
}

func parseSyslogFacility(name string) syslog.Priority {
	switch strings.ToLower(name) {
	case "local0":
		return syslog.LOG_LOCAL0
	case "local1":
		return syslog.LOG_LOCAL1
	case "local2":
		return syslog.LOG_LOCAL2
	case "local3":
		return syslog.LOG_LOCAL3
	case "local4":
		return syslog.LOG_LOCAL4
	case "local5":
		return syslog.LOG_LOCAL5
	case "local6":
		return syslog.LOG_LOCAL6
	case "local7":
		return syslog.LOG_LOCAL7
	default:
		return defaultSyslogFacility
	}
}

func parseSyslogSeverity(level string) syslog.Priority {
	switch strings.ToLower(level) {
	case "fatal":
		return syslog.LOG_EMERG
	case "critical":
		return syslog.LOG_CRIT
	case "error":
		return syslog.LOG_ERR
	case "warning":
		return syslog.LOG_WARNING
	case "info":
		return syslog.LOG_INFO
	case "debug":
		return syslog.LOG_DEBUG
	default:
		return defaultSyslogSeverity
	}
}

func isSyslogWriter(w io.Writer) bool {
	_, ok := w.(syslogIoWriterWrapper)
	return ok
}

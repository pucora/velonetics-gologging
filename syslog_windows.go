//go:build windows

package gologging

import (
	"fmt"
	"io"
)

type syslogPriority int

const (
	defaultSyslogFacility syslogPriority = 0
	defaultSyslogSeverity syslogPriority = 0
)

func newSyslogWriter(_ syslogPriority, _ syslogPriority, _ string) (io.Writer, error) {
	return nil, fmt.Errorf("syslog logging is not supported on Windows")
}

func parseSyslogFacility(_ string) syslogPriority {
	return defaultSyslogFacility
}

func parseSyslogSeverity(_ string) syslogPriority {
	return defaultSyslogSeverity
}

func isSyslogWriter(_ io.Writer) bool {
	return false
}

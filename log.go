package log

import (
	"fmt"
	"log"
	"log/syslog"
)

const priorityMask = 0x07

var (
	logger    = make([]*log.Logger, 0)
	syslogger = make([]*syslog.Writer, 0)
)

var syslogFunc = []func(*syslog.Writer, string) error{
	syslog.LOG_EMERG:   (*syslog.Writer).Emerg,
	syslog.LOG_ALERT:   (*syslog.Writer).Alert,
	syslog.LOG_CRIT:    (*syslog.Writer).Crit,
	syslog.LOG_ERR:     (*syslog.Writer).Err,
	syslog.LOG_WARNING: (*syslog.Writer).Warning,
	syslog.LOG_NOTICE:  (*syslog.Writer).Notice,
	syslog.LOG_INFO:    (*syslog.Writer).Info,
	syslog.LOG_DEBUG:   (*syslog.Writer).Debug,
}

// AddLogger adds a *log.Logger to log outputs
func AddLogger(l *log.Logger) {
	logger = append(logger, l)
}

// AddSyslog adds a *syslog.Writer to log outputs
func AddSyslog(l *syslog.Writer) {
	syslogger = append(syslogger, l)
}

func write(priority syslog.Priority, m string) error {
	for _, v := range logger {
		v.Print(m)
	}
	for _, v := range syslogger {
		if err := syslogFunc[priority&priorityMask](v, m); err != nil {
			return err
		}
	}
	return nil
}

func printf(priority syslog.Priority, format string, v ...interface{}) error {
	return write(priority, fmt.Sprintf(format, v...))
}

// Emerg logs a message with severety LOG_EMERG
func Emerg(format string, v ...interface{}) error {
	return printf(syslog.LOG_EMERG, format, v...)
}

// Alert logs a message with severety LOG_ALERT
func Alert(format string, v ...interface{}) error {
	return printf(syslog.LOG_ALERT, format, v...)
}

// Crit logs a message with severety LOG_CRIT
func Crit(format string, v ...interface{}) error {
	return printf(syslog.LOG_CRIT, format, v...)
}

// Err logs a message with severety LOG_ERR
func Err(format string, v ...interface{}) error {
	return printf(syslog.LOG_ERR, format, v...)
}

// Warning logs a message with severety LOG_WARNING
func Warning(format string, v ...interface{}) error {
	return printf(syslog.LOG_WARNING, format, v...)
}

// Notice logs a message with severety LOG_NOTICE
func Notice(format string, v ...interface{}) error {
	return printf(syslog.LOG_NOTICE, format, v...)
}

// Info logs a message with severety LOG_INFO
func Info(format string, v ...interface{}) error {
	return printf(syslog.LOG_INFO, format, v...)
}

// Debug logs a message with severety LOG_DEBUG
func Debug(format string, v ...interface{}) error {
	return printf(syslog.LOG_DEBUG, format, v...)
}

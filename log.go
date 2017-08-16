package log

import (
	"fmt"
	"log"
	"log/syslog"
)

const severityMask = 0x07

var (
	mask      = syslog.LOG_DEBUG
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
	if priority > mask {
		return nil
	}
	for _, v := range logger {
		v.Output(4, m)
	}
	for _, v := range syslogger {
		if err := syslogFunc[priority&severityMask](v, m); err != nil {
			return err
		}
	}
	return nil
}

func print(priority syslog.Priority, v ...interface{}) error {
	return write(priority, fmt.Sprint(v...))
}

func printf(priority syslog.Priority, format string, v ...interface{}) error {
	return write(priority, fmt.Sprintf(format, v...))
}

func println(priority syslog.Priority, v ...interface{}) error {
	return write(priority, fmt.Sprintln(v...))
}

// Mask causes future log messages with severety above priority
// not to be logged.
func Mask(priority syslog.Priority) {
	mask = priority
}

// Emerg logs a message with severety LOG_EMERG.
// Arguments are handled in the manner of fmt.Print.
func Emerg(v ...interface{}) error {
	return print(syslog.LOG_EMERG, v...)
}

// Alert logs a message with severety LOG_ALERT.
// Arguments are handled in the manner of fmt.Print.
func Alert(v ...interface{}) error {
	return print(syslog.LOG_ALERT, v...)
}

// Crit logs a message with severety LOG_CRIT.
// Arguments are handled in the manner of fmt.Print.
func Crit(v ...interface{}) error {
	return print(syslog.LOG_CRIT, v...)
}

// Err logs a message with severety LOG_ERR.
// Arguments are handled in the manner of fmt.Print.
func Err(v ...interface{}) error {
	return print(syslog.LOG_ERR, v...)
}

// Warning logs a message with severety LOG_WARNING.
// Arguments are handled in the manner of fmt.Print.
func Warning(v ...interface{}) error {
	return print(syslog.LOG_WARNING, v...)
}

// Notice logs a message with severety LOG_NOTICE.
// Arguments are handled in the manner of fmt.Print.
func Notice(v ...interface{}) error {
	return print(syslog.LOG_NOTICE, v...)
}

// Info logs a message with severety LOG_INFO.
// Arguments are handled in the manner of fmt.Print.
func Info(v ...interface{}) error {
	return print(syslog.LOG_INFO, v...)
}

// Debug logs a message with severety LOG_DEBUG.
// Arguments are handled in the manner of fmt.Print.
func Debug(v ...interface{}) error {
	return print(syslog.LOG_DEBUG, v...)
}

// Emergf logs a message with severety LOG_EMERG.
// Arguments are handled in the manner of fmt.Printf.
func Emergf(format string, v ...interface{}) error {
	return printf(syslog.LOG_EMERG, format, v...)
}

// Alertf logs a message with severety LOG_ALERT.
// Arguments are handled in the manner of fmt.Printf.
func Alertf(format string, v ...interface{}) error {
	return printf(syslog.LOG_ALERT, format, v...)
}

// Critf logs a message with severety LOG_CRIT.
// Arguments are handled in the manner of fmt.Printf.
func Critf(format string, v ...interface{}) error {
	return printf(syslog.LOG_CRIT, format, v...)
}

// Errf logs a message with severety LOG_ERR.
// Arguments are handled in the manner of fmt.Printf.
func Errf(format string, v ...interface{}) error {
	return printf(syslog.LOG_ERR, format, v...)
}

// Warningf logs a message with severety LOG_WARNING.
// Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, v ...interface{}) error {
	return printf(syslog.LOG_WARNING, format, v...)
}

// Noticef logs a message with severety LOG_NOTICE.
// Arguments are handled in the manner of fmt.Printf.
func Noticef(format string, v ...interface{}) error {
	return printf(syslog.LOG_NOTICE, format, v...)
}

// Infof logs a message with severety LOG_INFO.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) error {
	return printf(syslog.LOG_INFO, format, v...)
}

// Debugf logs a message with severety LOG_DEBUG.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...interface{}) error {
	return printf(syslog.LOG_DEBUG, format, v...)
}

// Emergln logs a message with severety LOG_EMERG.
// Arguments are handled in the manner of fmt.Println.
func Emergln(v ...interface{}) error {
	return println(syslog.LOG_EMERG, v...)
}

// Alertln logs a message with severety LOG_ALERT.
// Arguments are handled in the manner of fmt.Println.
func Alertln(v ...interface{}) error {
	return println(syslog.LOG_ALERT, v...)
}

// Critln logs a message with severety LOG_CRIT.
// Arguments are handled in the manner of fmt.Println.
func Critln(v ...interface{}) error {
	return println(syslog.LOG_CRIT, v...)
}

// Errln logs a message with severety LOG_ERR.
// Arguments are handled in the manner of fmt.Println.
func Errln(v ...interface{}) error {
	return println(syslog.LOG_ERR, v...)
}

// Warningln logs a message with severety LOG_WARNING.
// Arguments are handled in the manner of fmt.Println.
func Warningln(v ...interface{}) error {
	return println(syslog.LOG_WARNING, v...)
}

// Noticeln logs a message with severety LOG_NOTICE.
// Arguments are handled in the manner of fmt.Println.
func Noticeln(v ...interface{}) error {
	return println(syslog.LOG_NOTICE, v...)
}

// Infoln logs a message with severety LOG_INFO.
// Arguments are handled in the manner of fmt.Println.
func Infoln(v ...interface{}) error {
	return println(syslog.LOG_INFO, v...)
}

// Debugln logs a message with severety LOG_DEBUG.
// Arguments are handled in the manner of fmt.Println.
func Debugln(v ...interface{}) error {
	return println(syslog.LOG_DEBUG, v...)
}

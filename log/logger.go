package log

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Logger Serverities
const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

// Logger is logger
type Logger struct {
	Severity int
	Output   io.Writer
}

// Debug logs on DEBUG Serverity
func (l *Logger) Debug(format string, a ...interface{}) (n int, err error) {
	if l.Severity > DEBUG {
		return 0, nil
	}

	return fmt.Fprintf(l.Output, l.body("DEBUG", format), a...)
}

// Info logs on INFO Serverity
func (l *Logger) Info(format string, a ...interface{}) (n int, err error) {
	if l.Severity > INFO {
		return 0, nil
	}

	return fmt.Fprintf(l.Output, l.body("INFO", format), a...)
}

// Warn logs on WARN Serverity
func (l *Logger) Warn(format string, a ...interface{}) (n int, err error) {
	if l.Severity > WARN {
		return 0, nil
	}

	return fmt.Fprintf(l.Output, l.body("WARN", format), a...)
}

// Error logs on ERROR Serverity
func (l *Logger) Error(format string, a ...interface{}) (n int, err error) {
	if l.Severity > ERROR {
		return 0, nil
	}

	return fmt.Fprintf(l.Output, l.body("ERROR", format), a...)
}

// Fatal logs on FATAL Serverity
func (l *Logger) Fatal(format string, a ...interface{}) (n int, err error) {
	if l.Severity > FATAL {
		return 0, nil
	}

	return fmt.Fprintf(l.Output, l.body("FATAL", format), a...)
}

// NewLogger returns new Logger Instance.
func NewLogger(severity int, output io.Writer) *Logger {
	return &Logger{Severity: severity, Output: output}
}

// DefaultLogger returns DEBUG leveled logger outputting to STDOUT
func DefaultLogger() *Logger {
	return NewLogger(DEBUG, os.Stdout)
}

// private

func (l *Logger) now() string {
	return time.Now().String()
}

func (l *Logger) body(severity string, format string) string {
	return fmt.Sprintf("[%s: %s, PID: %d] %s\n", severity, l.now(), os.Getpid(), format)
}

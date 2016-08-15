package log

import "io"

// Usage
// l := log.NewAsyncLogger(log.DEBUG, ioutil.Discard)
// go l.Run()
// defer l.Stop()
// l.Info <- "Hello, World"

// AsyncLogger is logger that not locking
type AsyncLogger struct {
	Debug chan string
	Info  chan string
	Warn  chan string
	Error chan string
	Fatal chan string

	logger *Logger
	exit   chan bool
}

// Run makes AsyncLogger run
func (l *AsyncLogger) Run() {
	for {
		select {
		case message := <-l.Debug:
			l.logger.Debug(message)
		case message := <-l.Info:
			l.logger.Info(message)
		case message := <-l.Warn:
			l.logger.Warn(message)
		case message := <-l.Error:
			l.logger.Error(message)
		case message := <-l.Fatal:
			l.logger.Fatal(message)
		case <-l.exit:
			break
		}
	}
}

// Stop makes AsyncLogger stop
func (l *AsyncLogger) Stop() {
	l.exit <- true
}

// NewAsyncLogger returns new AsyncLogger
func NewAsyncLogger(severity int, output io.Writer) *AsyncLogger {
	return &AsyncLogger{
		Debug:  make(chan string),
		Info:   make(chan string),
		Warn:   make(chan string),
		Error:  make(chan string),
		Fatal:  make(chan string),
		logger: NewLogger(severity, output),
		exit:   make(chan bool),
	}
}

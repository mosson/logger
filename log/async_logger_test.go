package log

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestAsyncLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	l := NewAsyncLogger(DEBUG, buf)
	go l.Run()

	l.Info <- "Hello, World"

	l.Stop()

	if !bytes.Contains(buf.Bytes(), []byte("INFO:")) {
		t.Errorf("must contain 'INFO:', But actual output is %s", buf.Bytes())
	}

	if !bytes.Contains(buf.Bytes(), []byte("Hello, World")) {
		t.Errorf("must contain 'Hello, World', But actual output is %s", buf.Bytes())
	}

	buf2 := &bytes.Buffer{}
	l2 := NewAsyncLogger(FATAL, buf)
	go l2.Run()

	l2.Warn <- "Hello, World"

	l2.Stop()

	if bytes.Contains(buf2.Bytes(), []byte("Hello, World")) {
		t.Errorf("must not contain 'Hello, World', But actual output is %s", buf.Bytes())
	}
}

func TestAsyncLoggerCoverage(t *testing.T) {
	l := NewAsyncLogger(DEBUG, ioutil.Discard)
	go l.Run()
	defer l.Stop()

	l.Debug <- "Hello, World"
	l.Error <- "Hello, World"
	l.Fatal <- "Hello, World"
}

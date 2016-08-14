package log

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestInitialize(t *testing.T) {
	DefaultLogger()
}

func TestSeverity(t *testing.T) {
	buf := &bytes.Buffer{}
	l := NewLogger(DEBUG, buf)

	l.Info("Hello, World")

	if !bytes.Contains(buf.Bytes(), []byte("INFO:")) {
		t.Errorf("must contain 'INFO:', But actual output is %s", buf.Bytes())
	}

	if !bytes.Contains(buf.Bytes(), []byte("Hello, World")) {
		t.Errorf("must contain 'Hello, World', But actual output is %s", buf.Bytes())
	}

	buf2 := &bytes.Buffer{}
	l2 := NewLogger(FATAL, buf)

	l2.Warn("Hello, World")

	if bytes.Contains(buf2.Bytes(), []byte("Hello, World")) {
		t.Errorf("must not contain 'Hello, World', But actual output is %s", buf.Bytes())
	}
}

func TestCoverage(t *testing.T) {
	l := NewLogger(DEBUG, ioutil.Discard)
	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
	l.Error("error")
	l.Fatal("fatal")

	l2 := NewLogger(FATAL, ioutil.Discard)
	l2.Debug("debug")
	l2.Info("info")
	l2.Warn("warn")
	l2.Error("error")
	l2.Fatal("fatal")
}

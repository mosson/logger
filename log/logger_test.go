package log

import (
	"bytes"
	"testing"
)

func TestLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	l := NewLogger(DEBUG, buf)

	l.Info("Hello, World")

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

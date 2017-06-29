package log

import (
	"bufio"
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestLogger_LogInfo(t *testing.T) {
	buf := &bytes.Buffer{}
	w := bufio.NewWriter(buf)
	log.SetOutput(w)
	l := &Logger{}

	l.LogInfo("test")
	err := w.Flush()
	if err != nil {
		t.Error(err)
	}
	expected := "[INFO] test"
	if !strings.Contains(buf.String(), expected) {
		t.Errorf("Expected %s to include %s", buf.String(), expected)
	}

	l.LogInfo("test", "multiple")
	err = w.Flush()
	if err != nil {
		t.Error(err)
	}
	expected = "[INFO] test multiple"
	if !strings.Contains(buf.String(), expected) {
		t.Errorf("Expected %s to include %s", buf.String(), expected)
	}
}

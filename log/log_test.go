package log

import (
	"os"
	"testing"
)

func TestSetLevel(t *testing.T) {
	SetLevel(ErrorLevel)
	if warnLog.Writer() == os.Stdout || errorLog.Writer() != os.Stdout {
		t.Fatal("failed to set warn level")
	}
	SetLevel(Disabled)
	if infoLog.Writer() == os.Stdout || errorLog.Writer() == os.Stdout {
		t.Fatal("failed to set log level")
	}
}

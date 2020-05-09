package fleabug

import (
	"bytes"
	"testing"
)

func TestDefaultOutput(t *testing.T) {
	testOutput := new(bytes.Buffer)
	// init := DefaultOutput()
	SetDefaultOutput(testOutput)
	Dump("wopwop")
	if DefaultOutput() != testOutput {
		t.Errorf("failed to SetOutput")
	}
}

package fleabug

import (
	"bytes"
	// "fmt"
	"testing"
)

func TestDefaultOutputInjection(t *testing.T) {
	testOutput := new(bytes.Buffer)
	testOutput.Reset()

	dumper := Dumper{
		out: testOutput,
	}
	dumper.dump("wopwop")
	if DefaultOutput() == testOutput {
		t.Errorf("failed to SetOutput")
	}
	if testOutput.Len() != 16 {
		t.Errorf("testOutput don't return right value")
	}
}

func TestDumpWithString(t *testing.T) {
	testOutput := new(bytes.Buffer)
	testOutput.Reset()
	dumper := Dumper{
		out: testOutput,
	}
	dumper.dump("wopwop")
	expectedDump := "wopwop (string)\n"
	resultDump := string(testOutput.String())
	if expectedDump != resultDump {
		t.Errorf("the message is not the expected [" + expectedDump + "]:" + resultDump)
	}
}

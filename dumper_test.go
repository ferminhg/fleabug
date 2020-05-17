package fleabug

import (
	"bytes"
	"io"
	// "fmt"
	"testing"
	"unsafe"
)

func dumperFactory(out io.Writer) Dumper {
	return Dumper{
		out: out,
	}
}

func TestDumpWithString(t *testing.T) {
	testOutput := new(bytes.Buffer)
	dumper := dumperFactory(testOutput)

	dumper.dumpData("wopwop")
	expectedDump := "wopwop    (string)\n"
	resultDump := string(testOutput.String())
	if expectedDump != resultDump {
		t.Errorf("the message is not the expected [" + expectedDump + "]:" + resultDump)
	}
	if DefaultOutput() == testOutput {
		t.Errorf("failed to SetOutput")
	}
	if testOutput.Len() != 19 {
		t.Errorf("testOutput don't return right value")
	}
}

func TestDumpWithInt(t *testing.T) {
	testOutput := new(bytes.Buffer)
	dumper := dumperFactory(testOutput)
	dumper.dumpData(1)
	expectedDump := "1    (int)\n"
	if expectedDump != testOutput.String() {
		t.Errorf("the message is not the expected: " + expectedDump)
	}
}

func TestDumpTraceNoOk(t *testing.T) {
	testOutput := new(bytes.Buffer)
	dumper := dumperFactory(testOutput)
	var u = uintptr(unsafe.Pointer(testOutput))
	
	dumper.dumpTrace(u, "", 0, false)
	if "" != testOutput.String() {
		t.Errorf("the trace message is not empty")
	}

}


func TestDumpDummy(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	Dump("wopwop")
	Dump(1, 2, 3)
}
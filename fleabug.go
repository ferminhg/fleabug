package fleabug

import (
	"io"
	"log"
	"os"
	"runtime"
)

var defaultDumper = newDumper()
var defaultOut = os.Stdout

// Dump prints given arguments with newline.
func Dump(a ...interface{}) (n int, err error) {
	totalSize := 0
	n, _ = defaultDumper.blocker()
	totalSize += n
	n, err = defaultDumper.dumpData(a...)
	if err != nil {
		log.Fatal(err)
	}
	totalSize += n

	n, err = defaultDumper.dumpTrace(runtime.Caller(1))
	if err != nil {
		log.Fatal(err)
	}
	totalSize += n
	n, _ = defaultDumper.blocker()
	totalSize += n
	return totalSize, nil
}

// DefaultOutput returns dumper's default output.
func DefaultOutput() io.Writer {
	return defaultDumper.Output()
}

//SetDefaultOutput set new output to dumper's default
func SetDefaultOutput(o io.Writer) {
	defaultDumper.SetOutput(o)
}

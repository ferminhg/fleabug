package fleabug

import (
	"io"
	"os"
)

var defaultDumper = newDumper()
var defaultOut = os.Stdout

// Dump prints given arguments with newline.
func Dump(a ...interface{}) (n int, err error) {
	return defaultDumper.dump(a...)
}

// DefaultOutput returns dumper's default output.
func DefaultOutput() io.Writer {
	return defaultDumper.Output()
}

//SetDefaultOutput set new output to dumper's default
func SetDefaultOutput(o io.Writer) {
	defaultDumper.SetOutput(o)
}

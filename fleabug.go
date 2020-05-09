package fleabug

import (
	"fmt"
	"io"
	"os"
)

var defaultDumper = newDumper()
var defaultOut = os.Stdout

// Dumper provide the configuration needed to prettify
type Dumper struct {
	out io.Writer
	//configs
}

func newDumper() *Dumper {
	return &Dumper{
		out: defaultOut,
	}
}

// dump print default configuration given arguments
func (pp *Dumper) dump(params ...interface{}) (n int, err error) {
	return fmt.Println(pp.out, "wopwop")
}

// Output returns dumper's output.
func (pp *Dumper) Output() io.Writer {
	return pp.out
}

// SetOutput returns dumper's output.
func (pp *Dumper) SetOutput(out io.Writer) {
	// pp.outLock.Lock()
	pp.out = out
	// pp.outLock.Unlock()
}

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

package fleabug

import (
	"fmt"
	"io"
)

// Dumper provide the configuration needed to prettify
type Dumper struct {
	out io.Writer
	//configs
}

// newDumper named constructor Dumper
func newDumper() *Dumper {
	return &Dumper{
		out: defaultOut,
	}
}

// dump print default configuration given arguments
func (pp *Dumper) dump(params ...interface{}) (n int, err error) {
	//TODO: lanzar esto como una rutina
	var logger string
	for _, object := range params {
		logger += fmt.Sprintf("%v", formatter(object))
	}
	// fmt.Println("\t\tdump: [" + logger + "]")
	return fmt.Fprintln(pp.out, logger)
}

// Output returns dumper's output.
func (pp *Dumper) Output() io.Writer {
	return pp.out
}

// SetOutput returns dumper's output.
func (pp *Dumper) SetOutput(out io.Writer) {
	pp.out = out
}

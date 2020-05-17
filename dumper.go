package fleabug

import (
	"fmt"
	"io"
	"runtime"
	"strings"
)

// Dumper provide the configuration needed to prettify
type Dumper struct {
	out io.Writer
	//TODO: add more configs
}

// newDumper named constructor Dumper
func newDumper() *Dumper {
	return &Dumper{
		out: defaultOut,

	}
}

// dumpData print default configuration given arguments
func (pp *Dumper) dumpData(params ...interface{}) (n int, err error) {
	var logger, newLine string
	for item, object := range params {
		if newLine = ""; item < len(params)-1 {
			newLine = "\n"
		}
		logger += fmt.Sprintf("%v"+newLine, formatter(object))
	}
	// fmt.Println("\t\tdump: [" + logger + "]")
	return fmt.Fprintln(pp.out, logger)
}
//dumpTrace  print trace
func (pp *Dumper) dumpTrace(pc uintptr, file string, line int, ok bool) (n int, err error) {
	var traceMessage string
	if ok {
		traceMessage = fmt.Sprintf("# Called from %s line #%d \n# func: %v\n",
			file, line, runtime.FuncForPC(pc).Name())
	}
	return fmt.Fprint(pp.out, traceMessage)
}

func (pp *Dumper) blocker() (n int, err error){
	return fmt.Fprintln(pp.out, strings.Repeat(string("-"), 60))
}

// Output returns dumper's output.
func (pp *Dumper) Output() io.Writer {
	return pp.out
}

// SetOutput returns dumper's output.
func (pp *Dumper) SetOutput(out io.Writer) {
	pp.out = out
}

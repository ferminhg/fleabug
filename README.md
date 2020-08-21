# FleaBug: GoLang Dumper

[![Go Report Card](https://goreportcard.com/badge/github.com/ferminhg/fleabug)](https://goreportcard.com/report/github.com/ferminhg/fleabug)
[![Maintainability](https://api.codeclimate.com/v1/badges/e381261a2673537e1439/maintainability)](https://codeclimate.com/github/ferminhg/fleabug/maintainability)

FleaBug provides an easy and pretty `fmt.Println` alternative to `Go Lang` projects.
It shows value plus type and also trace (file, line, and func)

## Quick Start

Add this import line to the file you're working in:
```Go
import "github.com/ferminhg/fleabug"
```

```go
fleabug.Dump("wopwop")
fleabug.Dump(1, 2, 3)
````

## Sample Dump Output
```go
------------------------------------------------------------
wopwop  (string)
# Called from $HOME/go/src/github.com/ferminhg/fleabug/dumper_test.go line #62 
# func: github.com/ferminhg/fleabug.TestDumpDummy
------------------------------------------------------------

------------------------------------------------------------
1   (int)
2   (int)
3   (int)
# Called from $HOME/go/src/github.com/ferminhg/fleabug/dumper_test.go line #63 
# func: github.com/ferminhg/fleabug.TestDumpDummy
------------------------------------------------------------
````

## Installation

```bash
$ go get -u github.com/ferminhg/fleabug
```

## Credits

* Fermin Hernandez ([@fermin](https://www.linkedin.com/in/ferminhdez/))

## License

Fleabug is released under the MIT License. See the bundled LICENSE file for details.

FleaBug: GoLang Dumper
=========================

FleaBug provides an easy and pretty `fmt.Println` alternative to `Go Lang` projects.
It show value plus type and also trace (file, line, and func)

```go
fleabug.Dump("wopwop")
fleabug.Dump(1, 2, 3)
````

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

## Credits

* Fermin Hernandez ([@fermin](https://www.linkedin.com/in/ferminhdez/))

## License

Fleabug is released under the MIT License. See the bundled LICENSE file for details.

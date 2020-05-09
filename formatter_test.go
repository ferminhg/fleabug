package fleabug

import "testing"

func TestFormatterWithBasicTypes(t *testing.T) {
	if (formatter("wopwop") != "wopwop\t(string)") {
		t.Errorf("Bad format with string")
	}
	if (formatter(true) != "true\t(bool)") {
		t.Errorf("Bad format with bool")
	}
	var value int64
	value = 4
	if (formatter(value) != "4\t(int64)") {
		t.Errorf("Bad format with bool")
	}
}

func _TestFormatterWithFunc(t *testing.T){
	if (formatter(func(){}) !="\t(func)" ) {
		t.Errorf("Bad format with func: " + formatter(func(){}))
	}
}
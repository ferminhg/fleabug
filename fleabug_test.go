package fleabug

import "testing"

func TestDummy(t *testing.T) {
	if 1 != 1 {
		t.Errorf("world was changed")
	}
}

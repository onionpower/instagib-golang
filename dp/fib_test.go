package dp

import "testing"

func TestFibRec(t *testing.T) {
	//0 1 1 2 3 5 8
	f := fibr(7)
	if f != 8 {
		t.Errorf("expected is %v, but actual is %v", 8, f)
	}
}

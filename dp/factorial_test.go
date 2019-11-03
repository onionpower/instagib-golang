package dp

import (
	"fmt"
	"testing"
)

func TestBottomsUp(t *testing.T) {
	f := factBottomsUp(4)
	if f != 24 {
		t.Error(fmt.Sprintf("expected: 24, actual: %v", f))
	}
}

func TestMemo(t *testing.T) {
	f := factMemo(4)
	if f != 24 {
		t.Error(fmt.Sprintf("expected: 24, actual: %v", f))
	}
}

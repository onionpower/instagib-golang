package leetcode

import "testing"

func TestBinSearch704(t *testing.T) {
	tt := []struct {
		a   []int
		t   int
		exp int
	}{
		{
			a:   []int{-1, 0, 3, 5, 9, 12},
			t:   9,
			exp: 4,
		},
		{
			a:   []int{-1, 0, 3, 5, 9, 12},
			t:   2,
			exp: -1,
		},
	}
	for _, te := range tt {
		a := search(te.a, te.t)
		if a != te.exp {
			t.Errorf("expected is %v, but actual is %v", te.exp, a)
		}
	}

}

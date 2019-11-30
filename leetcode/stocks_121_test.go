package leetcode

import "testing"

func TestMaxProfit121(t *testing.T) {
	assertMaxProfit121([]int{7,1,5,3,6,4}, 5, t)
	assertMaxProfit121([]int{7,6,4,3,1}, 0, t)
}

func assertMaxProfit121(in []int, exp int, t *testing.T) {
	act := maxProfit121(in)
	if act != exp {
		t.Errorf("expected is %v, but actual is %v", exp, act)
	}
}

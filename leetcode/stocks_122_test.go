package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	assertMaxProfit([]int{7, 1, 5, 3, 6, 4}, 7, t)
	assertMaxProfit([]int{1, 2, 3, 4, 5}, 4, t)
	assertMaxProfit([]int{7, 6, 4, 3, 1}, 0, t)
}

func assertMaxProfit(in []int, exp int, t *testing.T) {
	p := maxProfit(in)
	if p != exp {
		t.Errorf("actual is %v, but expected is %v", p, exp)
	}
}

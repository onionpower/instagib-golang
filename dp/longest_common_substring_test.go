package dp

import t "testing"

func TestLongestCommonSubstring(t *t.T) {
	lcsAssert(longestCommonSubstring, "abcdf", "eebcdg", "bcd", t)
}

func lcsAssert(lcs func(s1 string, s2 string) string, s1 string, s2 string, exp string, t *t.T) {
	act := lcs(s1, s2)
	if exp != act {
		t.Errorf("actual is %v, but expected is %v", act, exp)
	}
}

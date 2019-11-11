package dp

import t "testing"

func TestLongestCommonSubstring(t *t.T) {
	lcsAssert(longestCommonSubstring, "uabcdf", "ueebcdg", "ubcd", t)
}

func TestLongestCommonSubstringMemo(t *t.T) {
	lcsAssert(longestCommonSubstringMemo, "uabcdf", "ueebcdg", "ubcd", t)
}

func BenchmarkLcs(b *t.B) {
	for n := 0; n < b.N; n++ {
		longestCommonSubstring("uabcdf", "ueebcdg")
	}
}

func BenchmarkLcsMemo(b *t.B) {
	for n := 0; n < b.N; n++ {
		longestCommonSubstringMemo("uabcdf", "ueebcdg")
	}
}

func lcsAssert(lcs func(s1 string, s2 string) string, s1 string, s2 string, exp string, t *t.T) {
	act := lcs(s1, s2)
	if exp != act {
		t.Errorf("actual is %v, but expected is %v", act, exp)
	}
}

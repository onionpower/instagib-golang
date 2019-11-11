package dp

func longestCommonSubstring(s1, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 {
		return ""
	}

	if s1[len(s1)-1] == s2[len(s2)-1] {
		return longestCommonSubstring(s1[:len(s1)-1], s2[:len(s2)-1]) + string(s1[len(s1)-1])
	}

	sl := longestCommonSubstring(s1, s2[:len(s2)-1])
	sr := longestCommonSubstring(s1[:len(s1)-1], s2)
	if len(sl) > len(sr) {
		return sl
	}
	return sr
}

func longestCommonSubstringMemo(s1, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 {
		return ""
	}

	m := make([][]string, len(s1)+1)
	for i := range m {
		m[i] = make([]string, len(s2)+1)
		for j := range m[i] {
			m[i][j] = "_"
		}
	}

	return longestCommonSubsMemoI(s1, s2, m)
}

func longestCommonSubsMemoI(s1 string, s2 string, m [][]string) string {
	if len(s1) == 0 || len(s2) == 0 {
		return ""
	}

	memo := m[len(s1)][len(s2)]
	if memo != "_" {
		return memo
	}

	li := len(s1) - 1
	ri := len(s2) - 1
	if s1[li] == s2[ri] {
		m[li][ri] = longestCommonSubsMemoI(
			s1[:li], s2[:ri], m) + string(s2[ri])
		return m[len(s1)-1][len(s2)-1]
	}

	sl := m[len(s1)-1][len(s2)]
	if sl == "_" {
		sl = longestCommonSubsMemoI(s1[:len(s1)-1], s2, m)
		m[len(s1)-1][len(s2)] = sl
	}
	sr := m[len(s1)][len(s2)-1]
	if sr == "_" {
		sr = longestCommonSubsMemoI(s1, s2[:len(s2)-1], m)
		m[len(s1)][len(s2)-1] = sr
	}

	if len(sl) > len(sr) {
		return sl
	}

	return sr
}

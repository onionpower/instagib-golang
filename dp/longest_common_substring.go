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

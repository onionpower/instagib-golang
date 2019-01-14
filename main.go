package main

import (
	"fmt"
)

// TODO examine finding the shortest first
func longestCommonPrefix(strs []string) string {
	strs_len := len(strs)
	if strs_len == 0 {
		return ""
	}
	if strs_len == 1 {
		return strs[0]
	}

	var prefix []uint8
	// each char
	for ci := 0; ; ci++ {
		if ci >= len(strs[0]) {
			return string(prefix)
		}

		c := strs[0][ci]
		// each word
		for wi := 1; wi < strs_len; wi++ {
			w := strs[wi]
			if ci >= len(w) {
				return string(prefix)
			}
			if w[ci] != c {
				return string(prefix)
			}
		}
		prefix = append(prefix, c)
	}
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"abcd", "abc", "ab", "e"}))
}

package main

import (
	"fmt"
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	symbols := make(map[uint8]int)
	max := 0
	for l, r := 0, 0; r < len(s); r++ {
		char := s[r]
		ix, found := symbols[char]
		if found {
			l = maxInt(l, ix)
		}
		symbols[char] = r + 1
		max = maxInt(max, r-l+1)
	}
	return max
}

func main() {
	s := "asdf"
	assert(s, 4)
	s = "aadf"
	assert(s, 3)
	s = "asdd"
	assert(s, 3)
	s = "aadd"
	assert(s, 2)
	s = "pwwkew"
	assert(s, 3)
}

func assert(s string, expected int) {
	l := lengthOfLongestSubstring(s)
	report := fmt.Sprintf("%v\ngot: %v\nexpected: %v", s, l, expected)
	if l != expected {
		panic(report)
	}
	fmt.Println(report)
}

package main

import "testing"

func BenchmarkLongestLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring("pwwke")
	}
}

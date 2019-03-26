package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	return 0
}

func maxLeft() int {
	return 0
}

func main() {
	printSln("asdf")
}

func printSln(s string) {
	fmt.Println(lengthOfLongestSubstring(s))
}

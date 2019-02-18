package main

import (
	"fmt"
)

func mySqrt(x int) int {
	s := uint64(x)
	uix := uint64(x)
	for s*s > uix {
		s = (s + uix/s) / 2
	}
	return int(s)
}

func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(10))
}

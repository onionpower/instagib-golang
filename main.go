package main

import "fmt"

func plusOne(digits []int) []int {
	if digits == nil {
		return []int{1}
	}

	return plusOneInternal(digits)
}

func plusOneInternal(digits []int) []int {
	o := 1
	for i := len(digits) - 1; i >= 0; i-- {
		t := digits[i] + o
		digits[i] = t % 10
		o = int(t / 10)
	}

	if o > 0 {
		return append([]int{1}, digits...)
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 0}))
	fmt.Println(plusOne([]int{1, 0, 1}))
	fmt.Println(plusOne([]int{9, 9}))
}

package main

import "fmt"

func reverse(l []int) []int {
	if len(l) == 1 {
		return l
	}
	return append(reverse(l[1:]), l[0])
}

func main() {
	fmt.Printf("%v", reverse([]int{1, 2, 3, 4}))
}

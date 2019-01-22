package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	if nums == nil {
		return 0
	}

	start, end := 0, len(nums)-1
	for start <= end {
		if nums[start] == val {
			nums[end], nums[start] = nums[start], nums[end]
			end--
		} else {
			start++
		}
	}
	return start
}

func main() {
	printRemoved([]int{0, 1, 2, 2, 3, 0, 4, 2}[:], 2)
	printRemoved([]int{1, 3, 3, 4}[:], 3)
	printRemoved([]int{1, 3, 3, 4}[:], 5)
	printRemoved([]int{1, 3, 3, 4}[:], 0)
	printRemoved([]int{3, 3}[:], 3)
	printRemoved([]int{3}[:], 3)
}

func printRemoved(nums []int, val int) {
	fmt.Printf("%v\n%v", nums, val)
	fmt.Println()
	length := removeElement(nums[:], val)
	fmt.Printf("%v", nums[0:length])
	fmt.Println()
	fmt.Println()
}

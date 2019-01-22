package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] < target {
			return 1
		}
		return 0
	}
	pivot := len(nums) / 2
	pv := nums[pivot]
	if pv == target {
		return pivot
	}
	if pv > target {
		return searchInsert(nums[:pivot], target)
	}
	return pivot + searchInsert(nums[pivot:], target)
}

func main() {
	printSln([]int{1, 3, 5, 6}[:], 5)
	printSln([]int{1, 3, 5, 6}[:], 1)
	printSln([]int{1, 3, 5, 6}[:], 7)
	printSln([]int{1, 3, 5, 6}[:], 0)
	printSln([]int{1, 3, 5, 6}[:], 2) // TODO wrong
}

func printSln(nums []int, target int) {
	fmt.Printf("%v, target: %v\n", nums, target)
	fmt.Printf("supposed pos is %v\n", searchInsert(nums, target))
}

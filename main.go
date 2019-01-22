package main

import (
	"fmt"
	"sort"
)

func removeElement(nums []int, val int) int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] == val {
			return false
		}
		if nums[j] == val {
			return true
		}
		return nums[i] < nums[j]
	})
	newLen := len(nums) - 1
	for ; newLen >= 0; newLen-- {
		if nums[newLen] != val {
			break
		}
	}
	return newLen + 1
}

func main() {
	printRemoved([]int{0, 1, 2, 2, 3, 0, 4, 2}[:], 2)
	printRemoved([]int{1, 3, 3, 4}[:], 3)
	printRemoved([]int{1, 3, 3, 4}[:], 5)
	printRemoved([]int{1, 3, 3, 4}[:], 0)
	printRemoved([]int{3, 3}[:], 3)
}

func printRemoved(nums []int, val int) {
	fmt.Printf("%v\n%v", nums, val)
	fmt.Println()
	length := removeElement(nums[:], val)
	fmt.Printf("%v", nums[0:length])
	fmt.Println()
	fmt.Println()
}

package main

import "fmt"

func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}
	ix := 0
	for i := 1; i < len(nums); i++ {
		if nums[ix] != nums[i] {
			ix++
			nums[ix] = nums[i]
		}
	}
	return ix + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{1, 2, 3}))
	fmt.Println(removeDuplicates([]int{1, 2, 2, 3}))
	fmt.Println(removeDuplicates([]int{1, 2, 3, 4}))
	fmt.Println(removeDuplicates([]int{2, 2, 3, 4}))
	fmt.Println(removeDuplicates([]int{2, 2, 3, 3}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

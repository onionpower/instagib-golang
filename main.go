package main

import "fmt"

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	return removeDuplicatesInternal(&nums, 0, 1)
}

func removeDuplicatesInternal(p_nums *[]int, it int, next int) int {
	nums := *p_nums
	if next >= len(nums) {
		return 1
	}
	if nums[it] == nums[next] {
		return removeDuplicatesInternal(p_nums, it, next+1)
	}
	it++
	nums[it] = nums[next]
	return 1 + removeDuplicatesInternal(p_nums, it, next+1)
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

package main

import "fmt"

func rotate(nums []int, k int) {
	k %= len(nums)
	if k == 0 {
		return
	}
	buf := make([]int, k)
	copy(buf, nums[:k])
	for i, ci := 0, 0; i < len(nums); i++ {
		ni := (i + k) % len(nums)
		tmp := nums[ni]
		nums[ni] = buf[ci]
		buf[ci] = tmp
		ci++
		if ci == len(buf) {
			ci = 0
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

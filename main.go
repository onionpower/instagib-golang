package main

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return append(make([][]int, 0), nums)
	}
	lv := nums[0]
	r := permute(nums[1:])
	l := []int{lv}
	l = append(l, r[0]...)
	r = append(r)
	r[0] = append(r[0], lv)
	r = append(r, l)
	return r
}

func main() {
	nums := []int{1, 2, 3}
	p := permute(nums)

	fmt.Printf("%v", p)
}

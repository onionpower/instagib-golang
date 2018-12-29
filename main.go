package main

import "fmt"

func main() {
	nums := []int{1, 2}
	pivot := int(len(nums) / 2)
	l := nums[0:pivot]
	m := nums[pivot]
	r := nums[pivot+1:]
	fmt.Printf("%v %v %v", l, m, r)
}

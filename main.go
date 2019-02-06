package main

import (
	"fmt"
)

func maxOf2(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxOf3(a int, b int, c int) int {
	if a > b {
		return maxOf2(a, c)
	}
	return maxOf2(b, c)
}

func maxIntersSum(l int, r int, p int) int {
	s := maxOf2(l+p, p)
	s = maxOf2(s+r, s)
	return s
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return maxOf3(nums[0], nums[1], nums[0]+nums[1])
	}
	pivot := len(nums) / 2
	maxLeft := maxLeftSum(nums[0:pivot])
	maxRight := maxRightSum(nums[pivot+1:])
	pivotSum :=
		maxOf3(maxSubArray(nums[0:pivot]),
			maxSubArray(nums[pivot+1:]),
			maxIntersSum(maxLeft, maxRight, nums[pivot]))
	return pivotSum
}
func maxRightSum(ints []int) int {
	if len(ints) == 0 {
		return 0
	}
	m := ints[0]
	tm := m
	for i := 1; i < len(ints); i++ {
		tm += ints[i]
		if tm > m {
			m = tm
		}
	}
	return m
}

func maxLeftSum(ints []int) int {
	m := ints[len(ints)-1]
	tm := m
	for i := len(ints) - 2; i >= 0; i-- {
		tm += ints[i]
		if tm > m {
			m = tm
		}
	}
	return m
}

func main() {
	printSubArr([]int{2, 1, -3, 4, -1, 2, 1, -5, 4})
	printSubArr([]int{-2, -5, 6, -2, -3, 1, 5, -6})
	printSubArr([]int{1, 2, 3})
	printSubArr([]int{-2, -1})
	printSubArr([]int{2, 1})
}

func printSubArr(ints []int) {
	fmt.Println(maxSubArray(ints))
}

package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func permute(nums []int) [][]int {
	l := len(nums)
	p := factorial(l)
	fp := make([][]int, p)
	for i := 0; i < p; i++ {
		a := make([]int, l)
		a[0] = nums[i%l]
		fp[i] = a
	}

	it := 0
	for row := 1; row < p; row++ {
		for col := 0; col < l; col++ {
			var prev int
			if col == 0 {
				prev = fp[row-1][l-1]
			} else {
				prev = fp[row-1][col]
			}
			for ; ; it = (it + 1) % l {
				if prev != nums[it] {
					fp[row][col] = nums[it] // TODO place prev here
					break
				}
			}
		}
	}

	return fp
}

func main() {
	//vr := make([][]int, 0)
	//vr = append(vr, []int{1,2,3})
	//vr = append(vr, []int{1,2,3})
	//fmt.Printf("%v", vr)
	//fmt.Printf("%v", vr[0][1])
	nums := []int{1, 2, 3}
	p := permute(nums)
	fmt.Printf("%v", p)
}

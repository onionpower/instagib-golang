package main

import "fmt"

func permute(nums []int) [][]int {
	fp := make([][]int, 0)
	for i, v1 := range nums {
		p := make([]int, 0)
		p = append(p, v1)
		for j, v2 := range nums {
			if i == j {
				continue
			}
			p = append(p, v2)
		}
		fp = append(fp, p)
	}
	return fp
}

func main() {
	nums := []int{1, 2, 3}
	p := permute(nums)

	fmt.Printf("%v", p)
}

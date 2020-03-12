package leetcode

func permute(nums []int) [][]int {
	p := make([][]int, 0)
	perm(nums, &p, 0)
	return p
}

func perm(nums []int, p *[][]int, n int) {
	if n == len(nums) {
		a := make([]int, len(nums))
		copy(a, nums)
		*p = append(*p, a)
	} else {
		for i := n; i < len(nums); i++ {
			nums[i], nums[n] = nums[n], nums[i]
			perm(nums, p, n+1)
			nums[i], nums[n] = nums[n], nums[i]
		}
	}
}

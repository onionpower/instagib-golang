package leetcode

func maxProfit(prices []int) int {
	p := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			p += prices[i] - prices[i-1]
		}
	}
	return p
}

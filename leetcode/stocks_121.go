package leetcode

func maxProfit121(prices []int) int {
	min := 0
	max := 0
	p := 0
	t := 0
	for i := 0; i < len(prices); i++ {
		 if min < prices[i] {
		 	min = prices[i]
		 }
		 if max > prices[i] {
		 	max = prices[i]
		 }
		 pt := max - min
		 if pt > p {
		 	p = t
		 }
	}
	return p
}
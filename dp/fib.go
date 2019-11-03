package dp

func fibr(n int) int {
	if n <= 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	return fibr(n-1) + fibr(n-2)
}

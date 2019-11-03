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

func fibBotUp(n int) int {
	if n <= 1 {
		return 0
	}

	m := make([]int, n+1, n+1)
	m[2] = 1
	for i := 3; i <= n; i++ {
		m[i] = m[i-1] + m[i-2]
	}

	return m[n]
}

package dp

func factBottomsUp(n int) int {
	if n <= 1 {
		return 1
	}

	f := make([]int, n+1, n+1)
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] * i
	}

	return f[n]
}

func factMemo(n int) int {
	if n <= 1 {
		return 1
	}

	fm := make([]int, n+1, n+1)
	return factMemoF(fm, n)
}

func factMemoF(f []int, n int) int {
	if n == 1 {
		return 1
	}

	f[n] = n * factMemoF(f, n-1)
	return f[n]
}

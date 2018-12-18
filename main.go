package main

import "fmt"

func reverse(x int) int {
	neg := false
	if x < 0 {
		x *= -1
		neg = true
	}
	rem := 0
	r := 0
	for ; x > 0; x /= 10 {
		rem = x % 10
		r *= 10
		r += rem
	}

	if r > 2147483647 {
		return 0
	}

	if neg {
		r *= -1
	}

	return r
}

func main() {
	fmt.Println(reverse(123456))
	fmt.Println(reverse(-321))
	fmt.Println(reverse(120))
}

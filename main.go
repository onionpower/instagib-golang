package main

import (
	"fmt"
	_ "github.com/segmentio/kafka-go/lz4"
)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	return myPowInternal(x, n)
}

func myPowInternal(f float64, i int) float64 {
	var res float64 = 1
	for ; i > 0; i = i >> 1 {
		if i&1 == 1 {
			res *= f
		}

		f = f * f
	}
	return res
}

func main() {
	fmt.Println(myPow(4, 5))
	fmt.Println(myPow(4, -2))
}

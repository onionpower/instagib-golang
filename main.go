package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 || x != 0 && x%10 == 0 {
		return false
	}
	y := 0
	for ; x > y; x /= 10 {
		y *= 10
		y += x % 10
	}

	return x == y || x == y/10
}

func main() {
	fmt.Println(isPalindrome(-123))
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(1321))
	fmt.Println(isPalindrome(1021))
	fmt.Println(isPalindrome(101))
	fmt.Println(isPalindrome(1221))
	fmt.Println(isPalindrome(1001))
}

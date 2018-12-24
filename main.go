package main

import (
	"container/list"
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	l := list.New()
	digits := 0
	for ; x > 0; x /= 10 {
		l.PushBack(x % 10)
		digits++
	}
	ls := l.Front()
	rs := l.Back()
	for rs != ls {
		if rs.Value != ls.Value {
			return false
		}
		rs = rs.Prev()
		ls = ls.Next()
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(-123))
	fmt.Println(isPalindrome(1221))
	fmt.Println(isPalindrome(-1221))
	fmt.Println(isPalindrome(12321))
}

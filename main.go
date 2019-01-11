package main

import "fmt"

func romanToInt(s string) int {
	w := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	if s == "" {
		return 0
	}
	r := 0
	greatest := 0
	for i := len(s) - 1; i >= 0; i-- {
		cw := w[s[i]]
		if cw > greatest {
			r += cw
			greatest = cw
		} else if cw < greatest {
			r -= cw
		} else {
			r += cw
		}
	}
	return r
}

func main() {
	fmt.Println(romanToInt("IIIV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("IIIL"))
}

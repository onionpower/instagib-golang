package main

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

//todo make it with strconv

func charToNum(r uint8) (uint8, error) {
	if r < '0' || r > '9' {
		return 0, errors.New("")
	}
	return r - '0', nil
}

func addBinary(a string, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}

	less := len(a)
	if len(b) < len(a) {
		less = len(b)
	}

	var rem uint8 = 0
	res := make([]string, 0, len(a)+len(b))
	i := less - 1
	for ; i >= 0; i-- {
		// todo check the numers are valid
		l, err := charToNum(a[i])
		if err != nil {
			return ""
		}

		r, err := charToNum(b[i])
		if err != nil {
			return ""
		}

		sum := l + r + rem
		digit := sum % 2
		res = append(res, strconv.Itoa(int(digit)))
		rem = sum - digit
	}

	return strings.Join(res, "")
}

func main() {
	//fmt.Println(addBinary("000", "111"))
	//fmt.Println(addBinary("000", "101"))
	//fmt.Println(addBinary("", "101"))
	fmt.Println(addBinary("1", "101"))
}

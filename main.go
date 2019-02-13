package main

import (
	"fmt"
	"github.com/pkg/errors"

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

	return addBinaryInternal(a, b)
}

func addBinaryInternal(a string, b string) string {
	res := make([]uint8, 0, len(a)+len(b))
	var rem uint8 = 0
	bTail := len(b) - 1
	aTail := len(a) - 1

	for ; aTail >= 0 && bTail >= 0; aTail, bTail = aTail-1, bTail-1 {
		aDigit, err := charToNum(a[aTail])
		if err != nil {
			return ""
		}

		bDigit, err := charToNum(b[bTail])
		if err != nil {
			return ""
		}

		s := aDigit + bDigit + rem
		d := s % 2
		res = append(res, d)
		rem = s - d
	}

	for ; aTail >= 0; aTail-- {
		aDigit, err := charToNum(a[aTail])
		if err != nil {
			return ""
		}
		s := aDigit + rem
		d := s % 2
		res = append(res, d)
		rem = s - d
	}

	for ; bTail >= 0; bTail-- {
		bDigit, err := charToNum(b[bTail])
		if err != nil {
			return ""
		}
		s := bDigit + rem
		d := s % 2
		res = append(res, d)
		rem = s - d
	}

	for ; rem > 0; rem-- {
		res = append(res, 1)
	}

	sr := strings.Builder{}
	for i := len(res) - 1; i >= 0; i-- {
		sr.WriteString(fmt.Sprint(res[i]))
	}
	return sr.String()
}

func main() {
	//fmt.Println(addBinary("000", "101"))
	fmt.Println(addBinary("100", "11"))
	fmt.Println(addBinary("11", "111"))
	fmt.Println(addBinary("", "101"))
	fmt.Println(addBinary("1", "101"))
}

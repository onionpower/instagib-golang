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

func charsToNum(l uint8, r uint8) (uint8, uint8, error) {
	li, err := charToNum(l)
	if err != nil {
		return 0, 0, errors.New("convert err")
	}

	ri, err := charToNum(r)
	if err != nil {
		return 0, 0, errors.New("convert err")
	}

	return li, ri, nil
}

func addBinaryInternal2(a string, b string) string {
	ai, err := strconv.ParseInt(a, 2, 64)
	if err != nil {
		return ""
	}
	bi, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		return ""
	}
	return strconv.FormatInt(ai+bi, 2)
}

func addBinary(a string, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}

	return addBinaryInternal2(a, b)
}

func add(a uint8, b uint8, o uint8) (uint8, uint8) {
	s := a + b + o
	d := s % 2
	if s > 1 {
		return d, 1
	}
	return d, 0
}

func addBinaryInternal1(a string, b string) string {
	res := make([]uint8, 0, len(a)+len(b))
	var o uint8 = 0
	var d uint8
	bTail := len(b) - 1
	aTail := len(a) - 1

	for ; aTail >= 0 && bTail >= 0; aTail, bTail = aTail-1, bTail-1 {
		aDigit, bDigit, err := charsToNum(a[aTail], b[bTail])
		if err != nil {
			return ""
		}

		d, o = add(aDigit, bDigit, o)
		res = append(res, d)
	}

	for ; aTail >= 0; aTail-- {
		aDigit, err := charToNum(a[aTail])
		if err != nil {
			return ""
		}

		d, o = add(aDigit, 0, o)
		res = append(res, d)
	}

	for ; bTail >= 0; bTail-- {
		bDigit, err := charToNum(b[bTail])
		if err != nil {
			return ""
		}

		d, o = add(bDigit, 0, o)
		res = append(res, d)
	}

	if o > 0 {
		res = append(res, 1)
	}

	sr := strings.Builder{}
	for i := len(res) - 1; i >= 0; i-- {
		sr.WriteString(fmt.Sprint(res[i]))
	}
	return sr.String()
}

func main() {
	fmt.Println(addBinary("000", "101"))
	fmt.Println(addBinary("100", "11"))
	fmt.Println(addBinary("11", "111"))
	fmt.Println(addBinary("", "101"))
	fmt.Println(addBinary("1", "101"))
	fmt.Println(addBinary("1111", "1111"))
	fmt.Println(addBinary("11", "11010010"))
	fmt.Println(addBinary(
		"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
		"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
}

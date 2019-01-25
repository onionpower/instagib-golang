package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	symbols := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && symbols != 0 {
			return symbols
		}
		if s[i] != ' ' {
			symbols++
		}
	}
	return symbols
}

type Submission struct {
	ok     bool
	length int
}

type Res []Submission

var res Res

func (r Res) print() {
	ok := true
	rep := strings.Builder{}
	for i, v := range r {
		ok = ok && v.ok
		rep.WriteString(fmt.Sprintf("%v: %v is %v\n", i, v.length, v.ok))
	}
	fmt.Printf("%v\n", ok)
	fmt.Println(rep.String())
}

func main() {
	res = Res{}
	sln("a", 1)
	sln("asdf", 4)
	sln("", 0)
	sln(" ", 0)
	sln("   ", 0)
	sln("   ", 0)
	sln("as   ", 2)
	sln("a   asf   ", 3)
	sln("asd   asdf", 4)
	res.print()
}

func sln(s string, expected int) {
	l := lengthOfLastWord(s)
	res = append(res, Submission{l == expected, l})
}

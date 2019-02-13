package main

import "fmt"

func judgeCircle(moves string) bool {
	const u = 'U'
	const d = 'D'
	const r = 'R'
	const l = 'L'
	up := 0
	right := 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case u:
			up++
		case d:
			up--
		case r:
			right++
		case l:
			right--
		default:
			return false
		}
	}

	if up != 0 || right != 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println(judgeCircle("UUDD"))
	fmt.Println(judgeCircle("LL"))
	fmt.Println(judgeCircle("RLUURDDDLU"))
}

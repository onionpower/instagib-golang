package main

import "fmt"

func judgeCircle(moves string) bool {
	const u = 'U'
	const d = 'D'
	const r = 'R'
	const l = 'L'
	m := make(map[uint8]int, 2)
	m[u] = 0
	m[r] = 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case u:
			m[u]++
		case d:
			m[u]--
		case r:
			m[r]++
		case l:
			m[r]--
		default:
			return false
		}
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(judgeCircle("UUDD"))
	fmt.Println(judgeCircle("LL"))
	fmt.Println(judgeCircle("RLUURDDDLU"))
}

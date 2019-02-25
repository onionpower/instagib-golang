package main

import "fmt"

// x1 + v1 * y = x2 + v2 * y
// x1 + v1 * y - v2 * y = x2
// y * (v1 - v2) = x2 - x1
// y = (x2 - x1)/(v1 - v2)
func kangaroo(x1, v1, x2, v2 int32) string {
	if x1 < 0 || x2 < x1 || x2 > 10000 {
		return "NO"
	}
	if v1 < 1 || v1 > 10000 {
		return "NO"
	}
	if v2 < 1 || v2 > 10000 {
		return "NO"
	}
	return kangarooInternal(float64(x1), float64(v1), float64(x2), float64(v2))
}

func kangarooInternal(x1, v1, x2, v2 float64) string {
	steps := (x2 - x1) / (v1 - v2)
	if steps > 0 && steps <= 10000 && steps == float64(int32(steps)) {
		return "YES"
	}
	return "NO"
}

func main() {
	fmt.Println(kangaroo(0, 3, 4, 2))
	fmt.Println(kangaroo(0, 2, 5, 3))
	fmt.Println(kangaroo(21, 6, 47, 3))
}

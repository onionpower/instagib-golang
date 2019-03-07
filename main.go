package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	jewels := make(map[int32]bool, len(J))
	for _, j := range J {
		jewels[j] = true
	}

	own := 0
	for _, s := range S {
		if jewels[s] == true {
			own++
		}
	}

	return own
}

func main() {
	printSln("jak", "jjjjaaauUu")
}

func printSln(j string, s string) {
	fmt.Printf("%v %v %v\n", j, s, numJewelsInStones(j, s))
}

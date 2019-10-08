package main

import (
	"fmt"
	bp "instagob/boringpipelines"
)

func main() {
	ch := bp.Gen(1, 2, 3, 5)
	ch1 := bp.Sq(ch)
	ch2 := bp.Sq(ch)
	ch3 := bp.Sq(ch)
	mch := bp.Merge(ch1, ch2, ch3)
	for v := range mch {
		fmt.Println(v)
	}
}

package main

import (
	"fmt"
	bp "instagob/boringpipelines"
)

func main() {
	ch := bp.Sq(bp.Sq(bp.Gen(1, 2, 3)))
	for v := range ch {
		fmt.Println(v)
	}
}

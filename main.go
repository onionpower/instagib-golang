package main

import (
	"fmt"
	bc "instagob/basicconcurrency"
)

func main() {
	fmt.Println(10e2)
	ch := bc.FanIn(bc.BoringGen("me"), bc.BoringGen("you"))

	for i := 0; i < 25; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("done")
}

package main

import (
	"fmt"
	bc "instagob/basicconcurrency"
)

func main() {
	fmt.Println(10e2)
	ch := bc.BoringGen("me")
	ach := bc.BoringGen("you")
	for i := 0; i < 25; i++ {
		fmt.Println(<-ch)
		fmt.Println(<-ach)
	}
	fmt.Println("done")
}

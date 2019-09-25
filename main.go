package main

import (
	"fmt"
	bc "instagob/basicconcurrency"
	"math/rand"
	"time"
)

func main() {
	ch1 := bc.BoringGen("you", nil)
	ch2 := bc.BoringGen("me", nil)
	ch := bc.FanIn(ch1, ch2)
	done := time.After(time.Duration(rand.Intn(2e3)+3e3) * time.Millisecond)
	for {
		select {
		case v1 := <-ch:
			fmt.Println(v1)
		case <-done:
			fmt.Println("done")
			return
		}
	}
}

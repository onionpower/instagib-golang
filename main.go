package main

import (
	"fmt"
	bc "instagob/basicconcurrency"
	"math/rand"
	"time"
)

func main() {
	wait1 := make(chan bool)
	ch1 := bc.BoringGen("you", wait1)
	wait2 := make(chan bool)
	ch2 := bc.BoringGen("me", wait2)
	done := time.After(time.Duration(rand.Intn(2e3)+3e3) * time.Millisecond)
	for {
		select {
		case v1 := <-ch1:
			fmt.Println(v1)
			wait1 <- true
		case v2 := <-ch2:
			fmt.Println(v2)
			wait2 <- true
		case <-done:
			fmt.Println("done")
			return
		}
	}
}

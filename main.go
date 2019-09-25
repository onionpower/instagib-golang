package main

import (
	"fmt"
	bc "instagob/basicconcurrency"
	"math/rand"
	"time"
)

func main() {
	q := make(chan string)
	ch := bc.BoringGen("you", q)
	to := time.After(time.Duration(rand.Intn(2e3)+3e3) * time.Millisecond)
	for {
		select {
		case v1 := <-ch:
			fmt.Println(v1)
		case m := <-q:
			fmt.Println(m)
			return
		case <-to:
			fmt.Println("done")
			return
		}
	}
}

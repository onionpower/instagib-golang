package basicconcurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func Boring(msg string, ch chan<- string) {
	for i := 0; ; i++ {
		dur := time.Duration(rand.Intn(1e3)) * time.Millisecond
		ch <- fmt.Sprintln(msg, i, dur)
		time.Sleep(dur)
	}
}

func Listen(ch chan string) {
	for {
		fmt.Println(<-ch)
	}
}

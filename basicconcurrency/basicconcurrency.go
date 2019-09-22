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

func BoringGen(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			dur := time.Duration(rand.Intn(1e2)) * time.Millisecond
			ch <- fmt.Sprintf("%v boring %v", msg, dur)
			time.Sleep(dur)
		}
	}()
	return ch
}

func FanIn(ch1, ch2 <-chan string) chan string {
	ch := make(chan string)
	go func() {
		for {
			ch <- <-ch1
		}
	}()
	go func() {
		for {
			ch <- <-ch2
		}
	}()
	return ch
}

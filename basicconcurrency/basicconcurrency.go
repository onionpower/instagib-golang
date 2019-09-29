package basicconcurrency

import (
	"fmt"
	"math/rand"
	"time"
)

type Msg struct {
	str  string
	wait chan bool
}

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

func BoringGen(msg string, quit chan string) <-chan Msg {
	ch := make(chan Msg)
	go func() {
		st := time.Now()
		for {
			dur := time.Duration(rand.Intn(1e2)) * time.Millisecond
			ch <- Msg{
				fmt.Sprintf("%v boring %v", msg, dur),
				nil,
			}
			if time.Since(st) > 3*time.Second {
				quit <- "go away"
				return
			}
			time.Sleep(dur)
		}
	}()
	return ch
}

func FanIn(ch1, ch2 <-chan Msg) chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case msg := <-ch1:
				ch <- msg.str
			case msg := <-ch2:
				ch <- msg.str
			}
		}
	}()
	return ch
}

func DaisyChan(n int) {
	leftmost := make(chan int)
	l := leftmost
	r := leftmost
	for i := 0; i < n; i++ {
		r = make(chan int)
		go f(l, r)
		l = r
	}
	go func(c chan int) {
		c <- 1
	}(r)
	fmt.Println(<-leftmost)
}

func f(l, r chan int) {
	l <- 1 + <-r
}

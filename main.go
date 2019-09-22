package main

import (
	"fmt"
	bc "instagob/basicconcurrency"
	"time"
)

func main() {
	ch := make(chan string)
	go bc.Boring("boring", ch)
	go bc.Listen(ch)
	fmt.Println("listening")
	time.Sleep(time.Duration(2) * time.Second)
	close(ch)
	fmt.Println("i'm done")
}

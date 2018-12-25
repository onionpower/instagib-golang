package main

import (
	"fmt"
	"sync"
)

type A struct {
	Value string
}

func write(wg *sync.WaitGroup, val string, ch chan<- A) {
	defer wg.Done()
	fmt.Println(val, "will be written")
	ch <- A{Value: val}
	fmt.Println(val, "was written")
	fmt.Println(len(ch))
}

func read(ch <-chan A) {
	fmt.Println(len(ch))
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	deps := make(chan A, 3)
	defer close(deps)
	go write(&wg, "A", deps)
	go write(&wg, "B", deps)
	go write(&wg, "C", deps)
	wg.Wait()
	go read(deps)
	fmt.Println("done")
	fmt.Scanln()
}

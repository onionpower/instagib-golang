package boringpipelines

import "sync"

func Gen(n ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, i := range n {
			out <- i
		}
		close(out)
	}()
	return out
}

func GenBufd(n ...int) <-chan int {
	out := make(chan int, len(n))
	for _, v := range n {
		out <- v
	}
	close(out)
	return out
}

func Sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out
}

func Merge(chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(chs))
	fanin := func(in <-chan int) {
		for v := range in {
			out <- v
		}
		wg.Done()
	}
	for _, ch := range chs {
		go fanin(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

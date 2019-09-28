package main

import (
	"fmt"
	"instagob/chrdp"
	"sync"
	"time"
)

func main() {
	fmt.Println(benchmark(5))
}

func benchmark(n int) time.Duration {
	times := make(chan time.Duration)
	wg := &sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			now := time.Now()
			chrdp.MakeScreenshot()
			times <- time.Since(now)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(times)
	}()
	avg := time.Duration(0)
	for t := range times {
		avg += t
	}

	return time.Duration(int64(avg) / int64(n))
}

package main

import (
	"fmt"
	"instagob/chrdp"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	fu, err := ioutil.ReadFile("./chrdp/fu.html")
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/fu.html", func(wr http.ResponseWriter, r *http.Request) {
		now := time.Now()
		wr.Write(fu)
		fmt.Println(time.Since(now))
	})

	http.HandleFunc(`/fu.jpg`, func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		b := chrdp.MakeScreenshot(`http://localhost:5008/fu.html`)
		w.Write(b)
		fmt.Println(time.Since(now))
	})

	if err := http.ListenAndServe(":5008", nil); err != nil {
		fmt.Println(err)
	}
}

func benchmark(n int, url string) time.Duration {
	times := make(chan time.Duration)
	wg := &sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			now := time.Now()
			chrdp.MakeScreenshot(url)
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

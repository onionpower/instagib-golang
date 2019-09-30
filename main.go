package main

import (
	"fmt"
	"html/template"
	"instagob/chrdp"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	fu, err := ioutil.ReadFile("./chrdp/label.html")
	if err != nil {
		fmt.Println(err)
	}
	fonts, err := ioutil.ReadFile("./chrdp/code128.ttf")

	http.HandleFunc("/fu.html", func(wr http.ResponseWriter, r *http.Request) {
		now := time.Now()
		wr.Write(fu)
		fmt.Println(time.Since(now))
	})

	http.HandleFunc("/code128.ttf", func(w http.ResponseWriter, r *http.Request) {
		w.Write(fonts)
	})

	http.HandleFunc(`/fu.jpg`, func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		b := chrdp.MakeScreenshot(`http://localhost:5008/fu.html`)
		w.Write(b)
		fmt.Println(time.Since(now))
	})

	http.HandleFunc(`/tfu.jpg`, func(w http.ResponseWriter, r *http.Request) {
		t := template.New(`label.html`)
		t, err := t.ParseFiles(`./chrdp/label.html`)
		if err != nil {
			fmt.Print(err)
		}
		l := chrdp.NewLabel()
		err = t.Execute(w, l)
		if err != nil {
			fmt.Print(err)
		}
	})

	if err := http.ListenAndServe(":5008", nil); err != nil {
		fmt.Println(err)
	}

	http.Handle("/chrdp/", http.StripPrefix("/chrdp/", http.FileServer(http.Dir("chrdp"))))
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

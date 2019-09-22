package fakefetcher

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher, m *syncMap) {
	defer m.Done()
	if depth <= 0 {
		return
	}
	m.Set(url)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if m.Get(u) {
			continue
		}
		m.Add()
		go Crawl(u, depth-1, fetcher, m)
	}
}

func startFetching() {
	m := syncMap{
		m:  make(map[string]bool, 0),
		wg: &sync.WaitGroup{},
		mx: &sync.Mutex{},
	}
	m.Add()
	go Crawl("https://golang.org/", 4, fetcher, &m)
	m.wg.Wait()

}

type fakeFetcher map[string]*fakeResult

type syncMap struct {
	m  map[string]bool
	mx *sync.Mutex
	wg *sync.WaitGroup
}

func (m *syncMap) Set(k string) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.m[k] = true
}

func (m *syncMap) Get(k string) bool {
	m.mx.Lock()
	m.mx.Unlock()
	_, ok := m.m[k]
	return ok
}

func (m *syncMap) Add() {
	m.wg.Add(1)
}

func (m *syncMap) Done() {
	m.wg.Done()
}

func (m *syncMap) Wait() {
	m.wg.Wait()
}

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

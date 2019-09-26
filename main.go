package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"instagob/htmltopdf"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	fffuuu []byte
)

func main() {
	var err error
	fffuuu, err = ioutil.ReadFile("./htmltopdf/fffuuu.html")
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/fffuuu.html", func(wr http.ResponseWriter, r *http.Request) {
		now := time.Now()
		wr.Write(fffuuu)
		fmt.Println(time.Since(now))
	})

	http.HandleFunc("/fffuuu.pdf", func(wr http.ResponseWriter, r *http.Request) {
		now := time.Now()
		b := htmltopdf.GenPdf("http://localhost:5008/fffuuu.html")
		wr.WriteHeader(200)
		wr.Header().Add("content-type", "application/pdf")
		wr.Write(b)
		fmt.Println(time.Since(now))
	})

	http.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(":5008", nil); err != nil {
		fmt.Println(err)
	}
}

// TODO
// template engine: replace {todo} things with correspendent values
// metrics: measure response time
// stress test: 10 rps, response time. 100 rps

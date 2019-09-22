package mapclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func shoot(errs []error, m map[string]int, ammo interface{}, url string, n int) {
	buf, err := json.Marshal(ammo)
	if err != nil {
		errs = append(errs, err)
	}
	for i := 0; i < n; i++ {
		resp := post(buf, url, errs)
		m[resp] = m[resp] + 1
	}
	fmt.Println(m)
}

func post(buf []byte, url string, errs []error) string {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(buf))
	if err != nil {
		errs = append(errs, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errs = append(errs, err)
	}
	return string(body)
}

package main

import (
	"fmt"
	d "instagob/imgdraw"
	"time"
)

func main() {
	now := time.Now()
	d.Draw(200, 200)
	fmt.Println(time.Since(now))
}

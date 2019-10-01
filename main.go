package main

import (
	"fmt"
	"instagob/structs"
)

func main() {
	m := structs.NewHMap()
	for s := 'a'; s <= 'z'; s++ {
		m.Add(string(s), int(s))
	}
	fmt.Println(m)
	fmt.Print(m.Get("b"))
}

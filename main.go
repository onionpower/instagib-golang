package main

import (
	"fmt"
	"instagob/structs"
)

func main() {
	m := structs.NewHMap()
	for s := 'a'; s <= 'z'; s++ {
		m.Set(string(s), int(s))
	}
	fmt.Println(m)
	fmt.Println(m.Get("b"))
	fmt.Println(m.Delete("b"))
	fmt.Println(m.Get("b"))
	m.Set("b", 3)
	fmt.Println(m)
	fmt.Println(m.Get("b"))
	m.Set("b", 4)
	fmt.Println(m)
}

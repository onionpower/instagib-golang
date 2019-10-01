package main

import (
	"fmt"
	"hash/fnv"
	"instagob/structs"
)

func main() {
	m := structs.NewHMap()
	m.Add("a", 1)
	fmt.Println(m)
}

func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

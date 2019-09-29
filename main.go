package main

import (
	"fmt"
	s "instagob/structs"
	"math/rand"
)

func main() {
	a := make([]int64, 0)
	for i := 0; i < 30; i++ {
		a = append(a, int64(i))
	}
	h := s.NewBheap(a...)
	fmt.Println(h)

	a = make([]int64, 0)
	for i := 0; i < 30; i++ {
		a = append(a, int64(rand.Intn(50)))
	}
	h = s.NewBheap(a...)
	fmt.Println(h)
}

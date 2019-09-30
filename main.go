package main

import (
	"fmt"
	s "instagob/structs"
)

func main() {
	a := make([]int64, 0, 6)
	for i := 17; i < 22; i++ {
		a = append(a, int64(i))
	}
	h := s.NewMinHeap(a...)
	fmt.Println(h)
	err := h.Push(1)
	err = h.Push(2)
	err = h.Push(29)
	err = h.Push(0)
	fmt.Println(h)
	h.Pop()
	fmt.Print(h)
	if err != nil {
		fmt.Print(err)
	}
}

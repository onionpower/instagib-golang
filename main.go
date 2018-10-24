package main

import "fmt"

type RingBuffer struct {
	a    []int
	head int
}

func (rb *RingBuffer) Init(size int) {
	rb.a = make([]int, size)
	rb.head = 0
}

func (rb *RingBuffer) Push(i int) {
	rb.a[rb.head] = i
	rb.head++
	if rb.head > len(rb.a)-1 {
		rb.head = 0
	}
}

func main() {
	rb := RingBuffer{}
	rb.Init(5)
	for i := 0; i < 12; i++ {
		rb.Push(i)
	}
	fmt.Println(rb)
}

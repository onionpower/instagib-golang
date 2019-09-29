package structs

import (
	"errors"
	"fmt"
)

type Bheap []int64

func NewBheap(ps ...int64) Bheap {
	a := make([]int64, 0, len(ps))
	for _, p := range ps {
		a = append(a, p)
	}
	return a
}

func (h Bheap) String() string {
	var s string
	ln := 1
	for i, v := range h {
		if (i+1)%ln == 0 {
			s += "\n"
			ln *= 2
		}
		s += fmt.Sprintf("%v ", v)
	}
	return s
}

func (h Bheap) Add(n int64) error {
	if cap(h) == len(h) {
		return errors.New("heap is full")
	}
	h = append(h, n)

	for i := len(h) - 1; i >= 0; i-- {

	}

	return nil
}

// 1 2 3 4 5
// 1
// 1 2

package structs

import (
	"errors"
	"fmt"
)

type MinHeap struct {
	a   []int64
	len int
}

func NewMinHeap(ps ...int64) *MinHeap {
	a := make([]int64, 0, len(ps))
	h := MinHeap{}
	for _, p := range ps {
		a = append(a, p)
		h.len++
	}
	h.a = a
	return &h
}

func (h *MinHeap) String() string {
	var s string
	ln := 1
	for i, v := range h.a {
		if (i+1)%ln == 0 {
			s += "\n"
			ln *= 2
		}
		s += fmt.Sprintf("%v ", v)
	}
	return s
}

func (h *MinHeap) Push(v int64) error {
	h.len++
	h.a = append(h.a, v)
	for curIx, ancIx := h.len-1, 0; curIx >= 1; curIx = ancIx {
		ancIx = (curIx - 1) / 2
		if h.a[curIx] >= h.a[ancIx] {
			return nil
		}
		h.a[curIx], h.a[ancIx] = h.a[ancIx], h.a[curIx]
	}

	return nil
}

func (h *MinHeap) Pop() (int64, error) {
	if len(h.a) == 0 {
		return 0, errors.New("heap is empty")
	}

	v := h.a[0]
	h.a = h.a[1:]
	return v, nil
}

// 0 1 2 3 4
// 1 2 3 4 5
// 1
// 2 3
// 4 5 6

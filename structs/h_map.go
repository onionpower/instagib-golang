package structs

import (
	"errors"
	"fmt"
	"hash/fnv"
)

var KeyNotFoundErr = errors.New("key not found")

type Bucket struct {
	Key string
	Val int
}

func (b Bucket) String() string {
	return fmt.Sprintf("{%v:%v}", b.Key, b.Val)
}

type HMap struct {
	bs [][]Bucket
}

func NewHMap() *HMap {
	return &HMap{
		bs: make([][]Bucket, 10, 10),
	}
}

func (m *HMap) Set(key string, val int) {
	// TODO resize
	ix := m.ix(key)
	if m.bs[ix] == nil {
		m.bs[ix] = make([]Bucket, 0)
	}
	b := m.bs[ix]
	for bx := 0; bx < len(b); bx++ {
		if b[bx].Key == key {
			b[bx].Val = val
			return
		}
	}
	ab := &m.bs[ix]
	*ab = append(*ab, Bucket{
		Key: key,
		Val: val,
	})
}

func (m *HMap) Get(key string) (val int, err error) {
	ab, err := m.getBucket(key)
	if err != nil {
		return 0, err
	}

	for _, b := range ab {
		if b.Key == key {
			return b.Val, nil
		}
	}

	return 0, KeyNotFoundErr
}

func (m *HMap) Delete(key string) error {
	ix := m.ix(key)
	if m.bs[ix] == nil {
		return KeyNotFoundErr
	}

	b := m.bs[ix]
	for bx := 0; bx < len(b); bx++ {
		if b[bx].Key == key {
			//if bx == len(ab) - 1 {
			//	ab = ab[:bx]
			//} else {
			ab := &m.bs[ix]
			*ab = append(b[:bx], b[bx+1:]...)
			//}
			return nil
		}
	}

	return KeyNotFoundErr
}

func (m *HMap) getBucket(key string) ([]Bucket, error) {
	ix := m.ix(key)
	ab := m.bs[ix]
	if ab == nil {
		return nil, KeyNotFoundErr
	}
	return ab, nil
}

func (m *HMap) String() string {
	s := ""
	for _, ab := range m.bs {
		if ab != nil {
			for _, b := range ab {
				s += b.String()
			}
		}
	}
	return s
}

func (m *HMap) hash(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (m *HMap) ix(key string) uint32 {
	return m.hash(key) % uint32(len(m.bs))
}

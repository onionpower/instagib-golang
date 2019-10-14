package structs

import (
	"errors"
	"fmt"
	"hash/fnv"
)

type KeyNotFoundErr error

func NewKeyNotFound(key string) KeyNotFoundErr {
	return errors.New(fmt.Sprintf("\"%v\" key not found", key))
}

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
	b := m.bs[ix]
	for bx := 0; bx < len(b); bx++ {
		if b[bx].Key == key {
			b[bx].Val = val
			return
		}
	}
	abs := &m.bs[ix]
	*abs = append(*abs, Bucket{
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

	return 0, NewKeyNotFound(key)
}

func (m *HMap) Delete(key string) error {
	ab, err := m.getBucket(key)
	if err != nil {
		return err
	}

	for bx, b := range ab {
		if b.Key == key {
			ab = append(ab[:bx], ab[bx+1])
			return nil
		}
	}

	return NewKeyNotFound(key)
}

func (m *HMap) getBucket(key string) ([]Bucket, error) {
	ix := m.ix(key)
	ab := m.bs[ix]
	if ab == nil {
		return nil, NewKeyNotFound(key)
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

package structs

import (
	"fmt"
	"hash/fnv"
)

type Bucket struct {
	Key  string
	Val  int
	Next *Bucket
}

func (b Bucket) String() string {
	return fmt.Sprintf("{%v:%v}", b.Key, b.Val)
}

type HMap struct {
	bs []*Bucket
}

func NewHMap() *HMap {
	return &HMap{
		bs: make([]*Bucket, 10, 10),
	}
}

func (m *HMap) Add(key string, val int) {
	// TODO resize
	ix := m.hash(key) % uint32(len(m.bs))
	b := m.bs[ix]
	for ; b == nil; b = b.Next {

	}
	b = &Bucket{
		Key: key,
		Val: val,
	}
}

func (m *HMap) hash(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (m *HMap) String() string {
	s := ""
	for _, b := range m.bs {
		if b != nil {
			s += b.String()
			for n := b.Next; n != nil; n = n.Next {
				s += n.String()
			}
		}
	}
	return s
}

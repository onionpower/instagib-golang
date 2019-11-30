package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func TestAddTwoNumEqualLen(t *testing.T) {
	//243 564

	l1 := toLinkedList([]int{2, 4, 3})
	l2 := toLinkedList([]int{5, 6, 4})
	hl3 := addTwoNumbers(l1, l2)
	l3a := []int{7, 0, 8}
	assertAddTwoNum(hl3, l3a, t)
}

func TestAddTwoNumFstIsEmpty(t *testing.T) {
	l2 := toLinkedList([]int{5, 6, 4})
	hl3 := addTwoNumbers(nil, l2)
	assertAddTwoNum(hl3, []int{5, 6, 4}, t)
}

func TestAddTwoNumFstLonger(t *testing.T) {
	//243 564

	l1 := toLinkedList([]int{2, 4, 3, 3})
	l2 := toLinkedList([]int{5, 6, 4})
	hl3 := addTwoNumbers(l1, l2)
	l3a := []int{7, 0, 8, 3}
	assertAddTwoNum(hl3, l3a, t)

	hl3 = addTwoNumbers(l2, l1)
	assertAddTwoNum(hl3, l3a, t)
}

func assertAddTwoNum(hl3 *ListNode, l3a []int, t *testing.T) {
	itl3 := hl3
	for _, v := range l3a {
		if itl3.Val != v {
			t.Errorf("actual is %v, but expected is %v", itl3.Val, v)
		}
		itl3 = itl3.Next
	}
	if itl3 != nil {
		t.Errorf("actual list %v is longer, than expected %v", linkedToString(hl3), l3a)
	}
	fmt.Println(linkedToString(hl3))
}

func TestToLinkedList(t *testing.T) {
	a := []int{2, 4, 3}
	l := toLinkedList(a)
	for i := 0;l != nil; l = l.Next {
		if a[i] != l.Val {
			t.Errorf("%v is expected to be equal to %v", l.Val, a[i])
		}
		i++
	}
}

func TestPrintLinked(t *testing.T){
	l := toLinkedList([]int{1, 2, 3})
	p := linkedToString(l)
	fmt.Println(p)
}

func toLinkedList(a []int) *ListNode {
	if len(a) < 1 {
		return nil
	}

	head := ListNode{}
	it := &head
	it.Val = a[0]
	for i := 1; i < len(a); i++ {
		n := &ListNode{}
		n.Val = a[i]
		it.Next = n
		it = it.Next
	}
	return &head
}

func linkedToString(l *ListNode) string {
	if l == nil {
		fmt.Println("nil")
		return "nil"
	}

	s := make([]string, 1)
	s = append(s, fmt.Sprintf("[ %v ", l.Val))
	for l = l.Next; l != nil; l = l.Next {
		s = append(s, fmt.Sprintf("%v ", l.Val))
	}
	s = append(s, fmt.Sprintf("]"))

	return strings.Join(s, "")
}

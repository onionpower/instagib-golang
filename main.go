package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1 == nil || l2 == nil {
		return nil
	}

	return mergeTwoListsInternal(l1, l2)
}

func mergeTwoListsInternal(l1 *ListNode, l2 *ListNode) *ListNode {
	lowest, left, right := nextLowest(l1, l2)
	res := &ListNode{lowest.Val, nil}
	it := res
	lowest, left, right = nextLowest(left, right)
	for ; lowest != nil; lowest, left, right = nextLowest(left, right) {
		it.Next = &ListNode{lowest.Val, nil}
		it = it.Next
	}
	return res
}

func nextLowest(l1 *ListNode, l2 *ListNode) (*ListNode, *ListNode, *ListNode) {
	if l1 == nil && l2 == nil {
		return nil, nil, nil
	}
	if l1 == nil {
		return l2, l1, l2.Next
	}
	if l2 == nil {
		return l1, l1.Next, l2
	}
	if l1.Val < l2.Val {
		return l1, l1.Next, l2
	}
	return l2, l1, l2.Next
}

func convToLl(ints []int) *ListNode {
	if ints == nil || len(ints) == 0 {
		return nil
	}
	first := &ListNode{ints[0], nil}
	it := first
	for i := 1; i < len(ints); i++ {
		it.Next = &ListNode{ints[i], nil}
		it = it.Next
	}
	return first
}

func llJoin(ll *ListNode) string {
	if ll == nil {
		return ""
	}

	str := strings.Builder{}
	for node := ll; node != nil; node = node.Next {
		str.WriteString(strconv.Itoa(node.Val))
	}
	return str.String()
}

func main() {
	l0 := convToLl([]int{1, 2, 3})
	fmt.Println(llJoin(l0))
	l1 := convToLl([]int{2, 4, 5})
	fmt.Println(llJoin(l1))
	res := mergeTwoLists(l0, l1)
	fmt.Println(llJoin(res))
}

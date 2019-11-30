package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h := &ListNode{}
	it := h
	m := 0
	if l1 != nil && l2 != nil {
		m = l1.Val + l2.Val
		it.Val = m % 10
		m /= 10

		l1 = l1.Next
		l2 = l2.Next
	}

	for ; l1 != nil && l2 != nil; it = it.Next {
		it.Next = &ListNode{}
		i := l1.Val + l2.Val + m
		it.Next.Val = i % 10
		m = i / 10

		l2 = l2.Next
		l1 = l1.Next
	}

	for ; l2 != nil; it = it.Next {
		it.Next = &ListNode{}
		it.Next.Val = l2.Val
		l2 = l2.Next
	}

	for ; l1 != nil; it = it.Next {
		it.Next = &ListNode{}
		it.Next.Val = l1.Val
		l1 = l1.Next
	}

	if m != 0 {
		it.Next = &ListNode{}
		it.Next.Val = m
	}
	return h
}


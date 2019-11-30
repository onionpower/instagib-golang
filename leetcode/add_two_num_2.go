package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h := &ListNode{}
	it := h
	m := 0
	for ; l1 != nil || l2 != nil; {
		l := 0
		if l1 != nil {
			l = l1.Val
		}
		r := 0
		if l2 != nil {
			r = l2.Val
		}
		n := &ListNode{}
		i := l + r + m
		n.Val = i % 10
		m = i / 10

		if l2 != nil {
			l2 = l2.Next
		}
		if l1 != nil {
			l1 = l1.Next
		}

		it.Next = n
		it = it.Next
	}

	if m != 0 {
		it.Next = &ListNode{}
		it.Next.Val = m
	}

	return h.Next
}

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycleConstSpace(head *ListNode) bool {
	for slow := head; slow != nil; slow = slow.Next {
		for fast, i := slow.Next, 0; fast != nil && i < 2; fast, i = fast.Next, i+1 {
			if fast == slow {
				return true
			}
			slow.Next = fast
		}
	}

	return false
}

func hasCycle(head *ListNode) bool {
	visited := make(map[*ListNode]struct{})
	for it := head; it != nil; it = it.Next {
		if _, ok := visited[it]; ok {
			return true
		}
		visited[it] = struct{}{}
	}
	return false
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	fmt.Println(hasCycleConstSpace(head))

	headL := &ListNode{
		Val: 4,
	}
	headC := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: headL,
			},
		},
	}
	headL.Next = headC
	fmt.Println(hasCycleConstSpace(headC))
}

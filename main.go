package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycleConstSpace(head *ListNode) bool {
	for it1 := head; it1 != nil; it1 = it1.Next {
		for it2 := it1.Next; it2 != nil; it2 = it2.Next {
			if it2 == it1 {
				return true
			}
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

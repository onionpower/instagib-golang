package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prev := head
	for it := head.Next; it != nil; it = it.Next {
		if it.Val == prev.Val {
			prev.Next = it.Next
			continue
		}
		prev = it
	}
	return head
}

func stringList(head *ListNode) string {
	if head == nil {
		return ""
	}

	b := strings.Builder{}
	for it := head; it != nil; it = it.Next {
		b.WriteString(fmt.Sprintf("%v, ", it.Val))
	}

	s := b.String()
	return s[:len(s)-2]
}

func toList(nums []int) *ListNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	head := &ListNode{
		Val: nums[0],
	}
	tail := head
	for i := 1; i < len(nums); i++ {
		tail.Next = &ListNode{
			Val: nums[i],
		}
		tail = tail.Next
	}

	return head
}

func main() {
	printSln(nil)
	printSln([]int{1})
	printSln([]int{1, 1, 1})
	printSln([]int{1, 2, 3, 4, 5, 6, 7})
	printSln([]int{1, 2, 3, 4, 5, 5, 5, 6, 7})
	printSln([]int{3, 3, 3, 145, 4, 4, 4})
	printSln([]int{6, 5, 5, 6, 5, 5})
}

func printSln(nums []int) {
	fmt.Println(nums)

	list := toList(nums)
	fmt.Println(stringList(list))

	cleanedList := deleteDuplicates(list)
	fmt.Println(stringList(cleanedList))
	fmt.Println()
}

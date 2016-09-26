package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//好久没刷题，还能一遍过... 开心
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	h := &ListNode{Next: head}
	thead := h
	for {
		if thead.Next != nil && thead.Next.Next != nil {
			tmp := thead.Next
			thead.Next = thead.Next.Next
			if thead.Next.Next != nil {
				tmp.Next = thead.Next.Next
			} else {
				tmp.Next = nil
			}
			thead.Next.Next = tmp
			if tmp.Next != nil {
				thead = thead.Next.Next
			} else {
				break
			}
		} else {
			break
		}
	}
	return h.Next
}

func main() {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2, Next: a}
	c := &ListNode{Val: 3, Next: b}
	d := &ListNode{Val: 4, Next: c}
	e := &ListNode{Val: 5, Next: d}
	f := &ListNode{Val: 6, Next: e}
	i := &ListNode{Val: 7, Next: f}
	j := &ListNode{Val: 8, Next: i}
	k := &ListNode{Val: 9, Next: j}

	result := swapPairs(k)
	// result := k

	tmp := 0
	for result.Next != nil {
		fmt.Println("result ", tmp, "val:", result.Val)
		result = result.Next
		tmp++
	}
	fmt.Println("result ", tmp, "val:", result.Val)
	fmt.Println("end")
}

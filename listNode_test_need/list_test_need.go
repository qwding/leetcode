package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return nil
}

func main() {
	a := GenList(4)
	b := GenList(5)
	c := GenList(7)
	d := GenList(8)

	fmt.Println("a ...")
	PrintList(a)
	fmt.Println("b ...")
	PrintList(b)
	fmt.Println("c ...")
	PrintList(c)
	fmt.Println("d ...")
	PrintList(d)
}

func GenList(length int) *ListNode {
	if length == 0 {
		return nil
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	head := &ListNode{}
	h := head
	for i := 0; i < length; i++ {
		h.Next = &ListNode{Val: r.Intn(15)}
		h = h.Next
	}
	return head.Next
}

func GenSortList(length int) *ListNode {
	if length == 0 {
		return nil
	}
	last := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	head := &ListNode{}
	h := head
	for i := 0; i < length; i++ {
		add := r.Intn(5)
		h.Next = &ListNode{Val: last + add}
		last = last + add
		h = h.Next
	}
	return head.Next
}

func PrintList(list *ListNode) {
	if list == nil {
		return
	}
	i := 0
	for {
		fmt.Println(i, "val:", list.Val)
		if list.Next == nil {
			return
		}
		list = list.Next
		i++
	}
}

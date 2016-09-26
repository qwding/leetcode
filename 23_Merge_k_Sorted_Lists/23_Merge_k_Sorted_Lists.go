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

	//这部分是坑
	k := 0
	for k != len(lists) {
		if lists[k] == nil {
			lists = append(lists[:k], lists[k+1:]...)
		} else {
			k++
		}
	}

	result := &ListNode{}
	h := result
	for len(lists) != 0 {
		min := lists[0].Val
		index := 0
		for i := 1; i < len(lists); i++ {
			if lists[i].Val < min {
				min = lists[i].Val
				index = i
			}
		}
		h.Next = lists[index]
		h = h.Next
		if lists[index].Next == nil {
			lists = append(lists[:index], lists[index+1:]...)
		} else {
			lists[index] = lists[index].Next
		}

	}

	return result.Next
}

func main() {
	a := GenSortList(4)
	b := GenSortList(5)
	c := GenSortList(7)
	d := GenSortList(8)

	fmt.Println("a ...")
	PrintList(a)
	fmt.Println("b ...")
	PrintList(b)
	fmt.Println("c ...")
	PrintList(c)
	fmt.Println("d ...")
	PrintList(d)

	lists := []*ListNode{nil, nil}
	// lists := []*ListNode{a, b, c, d}
	result := mergeKLists(lists)

	// result := a

	fmt.Println("result ...")
	PrintList(result)
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

package main

import (
// "fmt"
)

type LRUCache struct {
	Storage map[int]int
	Cap     int
	Length  int
	Head    *Node
	Tail    *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{}
	storage := make(map[int]int)
	lru.Storage = storage
	lru.Cap = capacity
	tail := &Node{}
	head := &Node{Next: tail}
	tail.Pre = head
	lru.Head = head
	lru.Tail = tail
	return lru
}

func (this *LRUCache) Get(key int) int {
	if rv, ok := this.Storage[key]; ok {
		this.MoveToHead(key)
		this.PrintAll()
		// println("get:", key, rv)
		return rv
	} else {
		// println("get:", key, -1)
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	// println("set", key, value)
	_, ok := this.Storage[key]
	if ok {
		this.Storage[key] = value
		this.MoveToHead(key)
	} else {
		if this.Length == this.Cap {
			this.DelTail()
		} else {
			this.Length += 1
		}
		this.Storage[key] = value
		node := &Node{Key: key}
		this.AddHead(node)
	}
	this.PrintAll()
}

func (this *LRUCache) PrintAll() {
	iter := this.Head
	// fmt.Printf("node ...")
	for iter != nil {
		// fmt.Printf("%d ", iter.Key)
		iter = iter.Next
	}
	// fmt.Printf("map %#v\n", this.Storage)
}

func (this *LRUCache) GetNode(key int) *Node {
	iterate := this.Head
	for iterate.Next != nil {
		iterate = iterate.Next
		if key == iterate.Key {
			return iterate
		}
	}
	return &Node{}
}

func (this *LRUCache) MoveToHead(key int) {
	node := this.GetNode(key)
	// fmt.Printf("MoveToHead %#v\n", node)
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	this.AddHead(node)
}

func (this *LRUCache) AddTail(node *Node) {
	pre := this.Tail.Pre
	pre.Next = node
	node.Pre = pre
	node.Next = this.Tail
	this.Tail.Pre = node
}

func (this *LRUCache) AddHead(node *Node) {
	// fmt.Printf("AddHead%#v\n", node)
	secode := this.Head.Next
	if secode == nil {
		node.Pre = this.Head
		this.Head.Next = node
		return
	} else {
		// fmt.Printf("add Head tail:%#v\n", this.Tail)
		secode.Pre = node
		node.Next = secode
		node.Pre = this.Head
		this.Head.Next = node
		// fmt.Printf("add2 Head tail:%#v\n", this.Tail)
	}
}

func (this *LRUCache) DelTail() int {
	// if this.Length == 1 {
	// 	return 0
	// }
	tail := this.Tail
	// fmt.Printf("tail: %#v\n", tail)
	// fmt.Printf("next: %#v\n", tail.Next)
	node := tail.Pre
	pre := node.Pre
	pre.Next = tail
	tail.Pre = pre

	delete(this.Storage, node.Key)
	// println("del", node.Key)
	return node.Key
}

type Node struct {
	Key  int
	Next *Node
	Pre  *Node
}

// func main() {
// 	lru := Constructor(2)
// 	lru.Put(1, 1)
// 	lru.Put(2, 2)
// 	println("1", lru.Get(1))
// 	lru.Put(3, 3)
// 	println("2", lru.Get(2))
// 	lru.Put(4, 4)
// 	println("3", lru.Get(3))
// 	println("4", lru.Get(4))
// }

func main() {
	lru := Constructor(2)
	lru.Get(2)
	lru.Put(2, 6)
	lru.Get(1)
	lru.Put(1, 5)
	lru.Put(1, 2)
	lru.Get(1)
	lru.Get(2)
}

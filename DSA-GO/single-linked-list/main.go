package main

import (
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
}
type LinkedList struct {
	first *Node
	last  *Node
	size  int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func NewNode(value interface{}) *Node {
	return &Node{value: value}
}

func (ll *LinkedList) AddLast(key interface{}) {
	node := NewNode(key)
	if ll.isEmpty() {
		ll.first = node
		ll.last = node
	} else {
		ll.last.next = node
		ll.last = node
	}
	ll.size++
}

func (ll *LinkedList) AddFirst(key interface{}) {
	node := NewNode(key)
	if ll.isEmpty() {
		ll.first = node
		ll.last = node
	} else {
		node.next = ll.first
		ll.first = node
	}
	ll.size++
}

func (ll *LinkedList) IndexOf(key interface{}) int {
	index := 0
	current := ll.first
	for current != nil {
		if current.value == key {
			return index
		}
		current = current.next
		index++
	}
	return -1
}

func (ll *LinkedList) Contains(key interface{}) bool {
	return ll.IndexOf(key) != -1
}

func (ll *LinkedList) RemoveFirst() {
	if ll.isEmpty() {
		panic("LIST IS EMPTY")
	}
	second := ll.first.next
	ll.first.next = nil
	ll.first = second
	ll.size--
}

func (ll *LinkedList) RemoveLast() {
	if ll.isEmpty() {
		panic("empty list")
	}
	if ll.first == ll.last {
		ll.first = nil
		ll.last = nil
	} else {
		previous := ll.getPrevious(ll.last)
		ll.last = previous
		ll.last.next = nil
	}
	ll.size--
}
func (ll *LinkedList) SizeOf() int {
	return ll.size
}

func (ll *LinkedList) ReverseList() {
	prev := ll.first
	curr := ll.first.next
	for curr != nil {
		n := curr.next
		curr.next = prev
	}
}

func (ll *LinkedList) getPrevious(node *Node) *Node {
	current := ll.first

	for current != nil {
		if current.next == node {
			return current
		}
		current = current.next
	}
	return nil
}
func (ll *LinkedList) isEmpty() bool {
	return ll.first == nil
}

func main() {
	linklist := NewLinkedList()
	fmt.Println(linklist.SizeOf())
	//linklist.AddLast(2)
	//linklist.AddLast(5)
	//linklist.AddLast(7)
	linklist.AddFirst(56)
	//linklist.AddFirst(96)
	fmt.Println(linklist.SizeOf())
	//linklist.RemoveFirst()
	linklist.RemoveFirst()
	//linklist.RemoveLast()
	fmt.Println(linklist.SizeOf())
}

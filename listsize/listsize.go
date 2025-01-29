// Write a function ListSize that returns the number of elements
// in a linked list l.

package main

import (
	"fmt"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	var count int

	for l.Head != nil {
		count++
		l.Head = l.Head.Next
	}
	return count
}

func ListPushFront(l *List, data interface{}) {
	temp := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = temp
		l.Tail = temp
	} else {
		temp.Next = l.Head
		l.Head = temp
	}
}

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}

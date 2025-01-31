// Write a function ListPushFront that inserts a new element
// NodeL at the beginning of the list l while using the structure List

package main

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
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
	ListPushFront(link, "man")
	ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}

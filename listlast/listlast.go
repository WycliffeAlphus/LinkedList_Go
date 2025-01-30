// Write a function ListLast that returns
// the Data interface of the last element of a linked list l.

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

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}

	current := l.Head

	for current.Next != nil {
		current = current.Next
	}

	return current.Data
}

func ListPushBack(l *List, data interface{}) {
	temp := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = temp
		l.Tail = temp
	} else {
		l.Tail.Next = temp
		l.Tail = temp
	}
}

func main() {
	link := &List{}
	link2 := &List{}

	ListPushBack(link, "three")
	ListPushBack(link, 3)
	ListPushBack(link, "1")

	fmt.Println(ListLast(link))
	fmt.Println(ListLast(link2))
}

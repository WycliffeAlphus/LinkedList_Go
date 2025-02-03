// Write a function ListReverse that reverses the order
// of the elements of a given linked list l.

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

func ListReverse(l *List) {
	var prev *NodeL

	l.Tail = l.Head

	current := l.Head

	for current != nil {
		next := current.Next

		current.Next = prev

		prev = current

		current = next
	}

	l.Head = prev
}

func main() {
	link := &List{}

	ListPushBack(link, 1)
	ListPushBack(link, 2)
	ListPushBack(link, 3)
	ListPushBack(link, 4)

	ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
}

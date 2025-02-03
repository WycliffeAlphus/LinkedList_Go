// Write a function ListAt that takes a pointer to
// the head of the list l and an int pos as parameters.
// This function should return the pointer to the NodeL
// in the position pos of the linked list l.
// In case of error the function should return nil.

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

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	curr := l
	for curr != nil && curr.Next != nil {
		curr = curr.Next
		count++
		if count == pos {
			return curr
		}
	}
	return nil
}
func main() {

	link := &List{}
	ListPushBack(link, "hello")
	ListPushBack(link, "how are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)

	fmt.Println(ListAt(link.Head, 3).Data)
	fmt.Println(ListAt(link.Head, 1).Data)
	fmt.Println(ListAt(link.Head, 7))

}

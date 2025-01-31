// Write a function ListClear
// that deletes all nodes from a linked list l.

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

func ListClear(l *List) {
	if l.Head != nil {
		l.Head = nil
	}
}

func PrintList(l *List) {
	link := l.Head

	for link != nil {
		fmt.Print(link.Data, "->")
		link = link.Next
	}
	fmt.Println(nil)
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

	ListPushBack(link, "I")
	ListPushBack(link, 1)
	ListPushBack(link, "something")
	ListPushBack(link, 2)

	fmt.Println("------list------")
	PrintList(link)
	ListClear(link)
	fmt.Println("------updated list------")
	PrintList(link)
}

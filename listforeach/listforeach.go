// Write a function ListForEach that applies a function given as argument to the data within each node of the list l.

// The function given as argument must have a pointer as argument: l *List

// Copy the functions Add2_node and Subtract3_node in the same file as the function ListForEach is defined.

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

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}

func ListForEach(l *List, f func(*NodeL)) {
	curr := l.Head

	for curr != nil {
		f(curr)
		curr = curr.Next
	}
}

func main() {
	link := &List{}

	ListPushBack(link, "1")
	ListPushBack(link, "2")
	ListPushBack(link, "3")
	ListPushBack(link, "5")

	ListForEach(link, Add2_node)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}
}

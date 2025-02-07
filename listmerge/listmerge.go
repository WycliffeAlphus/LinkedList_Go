// Write a function ListMerge that places elements of a list l2
// at the end of another list l1.

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

func ListMerge(l1 *List, l2 *List) {
	current := l1.Head

	var prev *NodeL

	if l1.Head == nil {
		l1.Head = l2.Head
	} else {

		for current != nil {
			prev = current
			current = current.Next
		}

		prev.Next = l2.Head

	}
}

func PrintList(l *List) {
	it := l.Head

	for it != nil {
		fmt.Print(it.Data, " -> ")

		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func ListPushBack(l *List, data interface{}) {
	temp := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = temp
		l.Tail = l.Head
	} else {

		l.Tail.Next = temp
		l.Tail = l.Tail.Next
	}
}

func main() {
	link := &List{}
	link2 := &List{}

	ListPushBack(link, "a")
	ListPushBack(link, "b")
	ListPushBack(link, "c")
	ListPushBack(link, "d")
	fmt.Println("-----first List------")
	PrintList(link)

	ListPushBack(link2, "e")
	ListPushBack(link2, "f")
	ListPushBack(link2, "g")
	ListPushBack(link2, "h")
	fmt.Println("-----second List------")
	PrintList(link2)

	fmt.Println("-----Merged List-----")
	ListMerge(link, link2)
	PrintList(link)
}

// Create a function named SortListInsert,
// which accepts a pointer to the head of a sorted linked list and an integer.

// The function should insert a new element into the list,
// with its Data set to the value of the integer.
// The element should be inserted so that the linked list remains sorted in ascending order by Data. The function should return the head of the list.

// Your function will only be tested with sorted linked lists.

package main

import "fmt"

type NodeI struct {
	Data int

	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	temp := &NodeI{Data: data_ref}

	if l == nil || l.Data >= data_ref {
		temp.Next = l
		return temp
	}

	curr := l
	var prev *NodeI

	for curr != nil && curr.Data < data_ref {
		prev = curr
		curr = curr.Next
	}

	prev.Next = temp
	temp.Next = curr

	return l
}

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *NodeI

	link = listPushBack(link, 1)
	link = listPushBack(link, 4)
	link = listPushBack(link, 9)

	PrintList(link)

	link = SortListInsert(link, -2)
	link = SortListInsert(link, 2)
	PrintList(link)
}

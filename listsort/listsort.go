// Write a function ListSort that sorts the
// nodes of a linked list by ascending order.

package main

import (
	"fmt"
)

type NodeI struct {
	Data int

	Next *NodeI
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

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}
	swapped := true

	for swapped {
		swapped = false
		curr := l

		for curr.Next != nil {

			if curr.Data > curr.Next.Data {
				curr.Data, curr.Next.Data = curr.Next.Data, curr.Data
				swapped = true
			}
			curr = curr.Next
		}
	}

	return l
}

func main() {
	var link *NodeI

	link = listPushBack(link, 5)
	link = listPushBack(link, 4)
	link = listPushBack(link, 3)
	link = listPushBack(link, 2)
	link = listPushBack(link, 1)

	PrintList(ListSort(link))
}

// Write a function SortedListMerge that merges two lists n1 and n2 in ascending order.

// During the tests n1 and n2 will already be initially sorted.

package main

import "fmt"

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

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil && n2 == nil {
		return nil
	}

	if n1 == nil {
		return n2
	}

	if n2 == nil {
		return n1
	}

	dummy := &NodeI{}

	tail := dummy

	for n1 != nil && n2 != nil {
		if n1.Data < n2.Data {
			tail.Next = n1
			n1 = n1.Next
		} else {
			tail.Next = n2
			n2 = n2.Next
		}
		tail = tail.Next
	}

	if n1 != nil {
		tail.Next = n1
	} else {
		tail.Next = n2
	}

	return dummy.Next
}

func main() {
	var link *NodeI
	var link2 *NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)

	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)

	PrintList(SortedListMerge(link2, link))
}

package linkedlist

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func New() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) Prepend(data int) {
	if list.head == nil {
		list.head = &Node{data: data}
		return
	}
	list.head = &Node{data: data, next: list.head}
}

func (list *LinkedList) String() string {
	if list.head == nil {
		return ""
	}

	var result string
	currentNode := list.head
	for currentNode != nil {
		result = result + fmt.Sprintf(" %v =>", currentNode.data)
		currentNode = currentNode.next
	}
	return result
}

func (list *LinkedList) DeleteWithValue(data int) {
	if list.head == nil {
		return
	}

	currentNode := list.head
	var previousNode *Node
	for currentNode != nil {
		if currentNode.data == data {
			if currentNode == list.head {
				list.head = currentNode.next
				return
			}
			previousNode.next = currentNode.next
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}
}

func (list *LinkedList) Contains(data int) bool {
	if list.head == nil {
		return false
	}

	currentNode := list.head

	for currentNode != nil {
		if currentNode.data == data {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

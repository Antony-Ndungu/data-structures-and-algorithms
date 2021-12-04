package bst

import (
	"fmt"
)

type node struct {
	data       int
	leftChild  *node
	rightChild *node
}

type bst struct {
	root *node
}

func New() *bst {
	return &bst{}
}

func (t *bst) Insert(data int) {
	newNode := &node{data: data}
	if t.root == nil {
		t.root = newNode
		return
	}

	currentNode := t.root
	for {
		if data < currentNode.data {
			if currentNode.leftChild == nil {
				currentNode.leftChild = newNode
				return
			}
			currentNode = currentNode.leftChild
		} else {
			if currentNode.rightChild == nil {
				currentNode.rightChild = newNode
				return
			}
			currentNode = currentNode.rightChild
		}
	}
}

func (t *bst) Traverse() {
	if t.root == nil {
		return
	}
	t.inOrderTraversal(t.root)
}

func (t *bst) Contains(data int) bool {
	if t.root == nil {
		return false
	}

	currentNode := t.root
	for currentNode != nil {
		if data < currentNode.data {
			currentNode = currentNode.leftChild
		} else if data > currentNode.data {
			currentNode = currentNode.rightChild
		} else {
			return true
		}
	}

	return false
}

func (t *bst) Delete(data int) {
	if t.root == nil {
		return
	}

	currentNode := t.root
	var parentNode *node
	for currentNode != nil {
		if data < currentNode.data {
			parentNode = currentNode
			currentNode = currentNode.leftChild
		} else if data > currentNode.data {
			parentNode = currentNode
			currentNode = currentNode.rightChild
		} else {
			if currentNode.leftChild == nil && currentNode.rightChild == nil {
				if parentNode == nil {
					t.root = nil
					return
				}
				if currentNode == parentNode.leftChild {
					parentNode.leftChild = nil
					return
				}
				parentNode.rightChild = nil
				return

			} else if currentNode.leftChild != nil && currentNode.rightChild == nil {
				if parentNode == nil {
					t.root = currentNode.leftChild
					return
				}

				if currentNode == parentNode.leftChild {
					parentNode.leftChild = currentNode.leftChild
					return
				}

				parentNode.rightChild = currentNode.leftChild
				return
			} else if currentNode.rightChild != nil && currentNode.leftChild == nil {
				if parentNode == nil {
					t.root = currentNode.rightChild
					return
				}

				if currentNode == parentNode.leftChild {
					parentNode.leftChild = currentNode.rightChild
					return
				}

				parentNode.rightChild = currentNode.rightChild
				return
			} else {
				predicessorParent, predicessor := t.getPredicessor(currentNode, currentNode.leftChild)
				predicessor.data, currentNode.data = currentNode.data, predicessor.data
				if predicessor == predicessorParent.leftChild {
					predicessorParent.leftChild = nil
					return
				}
				predicessorParent.rightChild = nil
				return
			}
		}
	}
}

func (t *bst) getPredicessor(parentNode, currentNode *node) (*node, *node) {
	for currentNode.rightChild != nil {
		parentNode = currentNode
		currentNode = currentNode.rightChild
	}
	return parentNode, currentNode
}

func (t *bst) inOrderTraversal(currentNode *node) {
	if currentNode.leftChild != nil {
		t.inOrderTraversal(currentNode.leftChild)
	}
	fmt.Printf("%v\t", currentNode.data)

	if currentNode.rightChild != nil {
		t.inOrderTraversal(currentNode.rightChild)
	}
}

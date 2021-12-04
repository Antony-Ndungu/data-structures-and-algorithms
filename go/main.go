package main

import (
	"fmt"

	"github.com/Antony-Ndungu/data-structures-and-algorithms/go/bst"
)

func main() {
	tree := bst.New()
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(0)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(5)
	tree.Traverse()
	fmt.Println("")
	tree.Delete(10)
	tree.Traverse()
	fmt.Println("")
}

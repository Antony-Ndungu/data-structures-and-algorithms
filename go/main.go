package main

import (
	"fmt"

	"github.com/Antony-Ndungu/data-structures-and-algorithms/go/linkedlist"
)

func main() {
	myLinkedList := linkedlist.New()
	fmt.Println(myLinkedList)
	myLinkedList.Prepend(5)
	fmt.Println(myLinkedList)
	myLinkedList.Prepend(6)
	fmt.Println(myLinkedList)
	myLinkedList.Prepend(2)
	fmt.Println(myLinkedList)
	fmt.Println(myLinkedList.Contains(5))
	myLinkedList.DeleteWithValue(5)
	fmt.Println(myLinkedList)
	fmt.Println(myLinkedList.Contains(5))
}

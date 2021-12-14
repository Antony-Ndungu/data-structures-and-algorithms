package main

import (
	"fmt"

	"github.com/Antony-Ndungu/data-structures-and-algorithms/go/heap"
)

func main() {
	nums := []int{27, 4, 7, 12, 9, 100}
	fmt.Println(nums)
	maxHeap := heap.NewMaxHeap()
	for _, num := range nums {
		maxHeap.Insert(num)
		fmt.Println(maxHeap.Elements)
	}
	fmt.Println(maxHeap.Extract())
	fmt.Println(maxHeap.Elements)
	fmt.Println(maxHeap.Extract())
	fmt.Println(maxHeap.Elements)
	fmt.Println(maxHeap.Extract())
	fmt.Println(maxHeap.Elements)
	fmt.Println(maxHeap.Extract())
	fmt.Println(maxHeap.Elements)
	fmt.Println(maxHeap.Extract())
	fmt.Println(maxHeap.Elements)
	fmt.Println(maxHeap.Extract())
	fmt.Println(maxHeap.Elements)
	fmt.Println(maxHeap.Extract())
	fmt.Println(maxHeap.Elements)
}

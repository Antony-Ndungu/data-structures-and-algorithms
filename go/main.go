package main

import (
	"fmt"

	"github.com/Antony-Ndungu/data-structures-and-algorithms/go/algorithms"
)

func main() {
	nums := []int{27, 4, 7, 12, 9, 100}
	sortedNums := algorithms.MergeSort(nums)
	fmt.Println(sortedNums)
}

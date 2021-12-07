package main

import (
	"fmt"

	"github.com/Antony-Ndungu/data-structures-and-algorithms/go/algorithms"
)

func main() {
	nums := []int{27, 4, 7, 12, 9}
	algorithms.SelectionSort(nums)
	fmt.Println(nums)
}

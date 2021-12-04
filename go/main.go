package main

import (
	"fmt"

	"github.com/Antony-Ndungu/data-structures-and-algorithms/go/hashmap"
)

func main() {
	myHashmap := hashmap.NewHashmap()
	myHashmap.Put("Antony", 1)
	myHashmap.Put("Ndungu", 2)
	value, ok := myHashmap.Get("Antony")
	if ok {
		fmt.Println(value)
	}
}

package main

import (
	"fmt"

	"github.com/Antony-Ndungu/data-structures-and-algorithms/go/queue"
)

func main() {
	myQueue := queue.New()
	myQueue.Enqueue(4)
	myQueue.Enqueue(3)
	myQueue.Enqueue(2)
	myQueue.Enqueue(1)
	for !myQueue.IsEmpty() {
		fmt.Println(myQueue.Dequeue())
	}
}

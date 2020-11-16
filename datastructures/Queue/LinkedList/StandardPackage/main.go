package main

import (
	"container/list"
	"fmt"
)

// https://yourbasic.org/golang/implement-fifo-queue/
// https://godoc.org/container/list

func main() {

	// container/list package implements a
	// doubly linked list which can be used as a queue

	queue := list.New()

	queue.PushBack("Hello ") // enqueue
	queue.PushBack("world!")

	for queue.Len() > 0 {
		e := queue.Front() // 1st element
		fmt.Print(e.Value)
		queue.Remove(e) // dequeue
	}

}

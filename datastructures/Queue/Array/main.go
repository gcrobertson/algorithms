package main

import "fmt"

// https://yourbasic.org/golang/implement-fifo-queue/

func main() {

	var queue []string

	queue = append(queue, "Hello ") // enqueue
	queue = append(queue, "world!")

	for len(queue) > 0 {
		fmt.Print(queue[0])
		queue[0] = ""     // **
		queue = queue[1:] // dequeue
	}

	// **
	// The memory allocated for the array is never returned.
	// For long-living queue you should probably use
	// a dynamic data structure, such as a linked list.
}

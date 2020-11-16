package main

import "fmt"

type node struct {
	val  string
	next *node
}

type list struct {
	head *node
	tail *node
}

var queue list

func (q list) enqueue(val string) {
	t := &node{
		val: val,
	}
	if q.head == nil {
		q.head = t
	} else {
		pt := q.tail
		pt.next = t

	}
	q.tail = t
	return
}

func main() {

	queue.enqueue("Hello ")
	queue.enqueue("World!")

	fmt.Print("starting for loop...")
	n := queue.head

	fmt.Print("value of n: ", n.val)

	// for n != nil {

	// 	fmt.Print(n.val + " ")

	// 	// dq := queue.dequeue()
	// 	// fmt.Print(dq.val + " ")
	// 	n = n.next
	// }

}

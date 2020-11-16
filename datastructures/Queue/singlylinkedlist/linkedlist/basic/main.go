package main

import "fmt"

type queue interface {
	isEmpty()          // O(1)
	peek() *node       // O(1)
	dequeue() *node    // O(1)
	enqueue(int) *node // O(1) *tail
	print()            // O(n)
}

type singlylinkedlist struct {
	root *node
}

type singlylinkedlistTail struct {
	singlylinkedlist
	tail *node
}

type node struct {
	value int
	next  *node
}

func (l *singlylinkedlist) isEmpty() bool {
	return l.root == nil
}

func (l *singlylinkedlist) peek() *node {
	return l.root
}

func (l *singlylinkedlist) print() {
	for n := l.root; n != nil; n = n.next {
		fmt.Printf("%d\t", n.value)
	}
	fmt.Println()
}

func (l *singlylinkedlist) dequeue() *node {
	t := l.root
	l.root = t.next
	return t
}

func (l *singlylinkedlist) enqueue(val int) *node {
	t := &node{
		value: val,
	}
	if l.root == nil {
		l.root = t
	} else {
		i := l.root
		for i.next != nil {
			i = i.next
		}
		i.next = t
	}
	return t
}

func (l *singlylinkedlistTail) enqueue(val int) *node {
	t := &node{
		value: val,
	}
	if l.root == nil {
		l.root = t
		l.tail = t
	} else {
		l.tail.next = t
		l.tail = t
	}
	return t
}

func main() {
	list1 := singlylinkedlist{}
	list2 := singlylinkedlistTail{}

	for _, v := range []int{5, 15, 155, 205, 399} {
		list1.enqueue(v)
		list2.enqueue(v)
	}
	list1.print()
	list2.print()

	for !list1.isEmpty() && !list2.isEmpty() {
		fmt.Printf("dequeued %d from list 1\n", list1.dequeue().value)
		fmt.Printf("dequeued %d from list 2\n", list2.dequeue().value)
	}
}

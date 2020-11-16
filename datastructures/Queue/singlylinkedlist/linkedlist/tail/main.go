package main

import "fmt"

type singlylinkedlist struct {
	root *node
	tail *node
}

type node struct {
	value int
	next  *node
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
		l.tail = t
	} else {
		l.tail.next = t
		l.tail = t
	}
	return t
}

func main() {
	singlylinkedlist := singlylinkedlist{}
	singlylinkedlist.enqueue(5)
	singlylinkedlist.enqueue(15)
	singlylinkedlist.enqueue(115)
	singlylinkedlist.print()
	for singlylinkedlist.root != nil {
		r := singlylinkedlist.dequeue()
		fmt.Printf("removed %d\t", r.value)
	}
}

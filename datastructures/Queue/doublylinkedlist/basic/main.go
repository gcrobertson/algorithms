package main

import (
	"fmt"
)

type queue interface {
	isEmpty() //	O(1)
	peek()    //	O(1)
	enqueue() //	O(1) if last node known else O(n)
	dequeue() //	O(1)
	print()   //	O(n)
}

type singlynode struct {
	value int
	next  *singlynode
}

type doublynode struct {
	value int
	next  *doublynode
	prev  *doublynode
}

type singlylinkedlist struct {
	root *singlynode
}

type doublylinkedlist struct {
	root *doublynode
}

type doublylinkedcircularlist struct {
	doublylinkedlist
	len int
}

type doublylinkedtailedlist struct {
	doublylinkedlist
	tail *doublynode
}

func (l *singlylinkedlist) isEmpty() bool {
	return l.root == nil
}

func (l *doublylinkedlist) isEmpty() bool {
	return l.root == nil
}

// superfluous if dequeue sets l.root to nil when c.len == 0
func (c *doublylinkedcircularlist) isEmpty() bool {
	return c.len == 0
}

func (l *singlylinkedlist) print() {
	for n := l.root; n != nil; n = n.next {
		fmt.Printf("%d\t", n.value)
	}
	fmt.Println()
}

func (l *doublylinkedlist) print() {
	for n := l.root; n != nil; n = n.next {
		fmt.Printf("%v\t", n.value)
	}
	fmt.Println()
}

func (c *doublylinkedcircularlist) print() error {
	n := c.root
	for i := c.len; i > 0; i-- {
		fmt.Printf("%v\t", n.value)
		n = n.next
	}
	fmt.Println()
	return nil
}

func (l *singlylinkedlist) peek() *singlynode {
	if l.root == nil {
		return nil
	}
	return l.root
}

func (l *doublylinkedlist) peek() *doublynode {
	if l.root == nil {
		return nil
	}
	return l.root
}

func (l *singlylinkedlist) dequeue() *singlynode {
	p := l.root
	l.root = p.next
	p.next = nil
	return p
}

func (l *doublylinkedlist) dequeue() *doublynode {
	p := l.root
	l.root = l.root.next
	p.next, p.prev = nil, nil
	return p
}

func (c *doublylinkedcircularlist) dequeue() *doublynode {
	p := c.root
	c.root = c.root.next
	c.root.prev = p.prev
	p.prev.next = c.root
	p.next, p.prev = nil, nil
	c.len--
	if c.len == 0 {
		c.root = nil
	}
	return p
}

func (l *singlylinkedlist) enqueue(val int) *singlynode {
	n := &singlynode{value: val}
	if l.root == nil {
		l.root = n
		return l.root
	}
	p := l.root
	for p.next != nil {
		p = p.next
	}
	p.next = n
	return n
}

func (l *doublylinkedtailedlist) enqueue(val int) *doublynode {
	p := &doublynode{value: val}
	if l.root == nil {
		l.root = p
	} else {
		l.tail.next = p
		p.prev = l.tail
	}
	l.tail = p
	return p
}

func (c *doublylinkedcircularlist) enqueue(val int) *doublynode {
	p := &doublynode{
		value: val,
	}
	if c.root == nil {
		c.root = p
		c.root.prev, c.root.next = p, p
	} else {
		c.root.prev.next = p
		p.prev = c.root.prev
		p.next = c.root
		c.root.prev = p
	}
	c.len++
	return p
}

func main() {

	singlylinkedlist := singlylinkedlist{}
	doublylinkedtailedlist := doublylinkedtailedlist{}
	doublylinkedcircularlist := doublylinkedcircularlist{}

	for _, v := range []int{5, 10, 15, 20, 25, 30, 35} {
		singlylinkedlist.enqueue(v)
		doublylinkedtailedlist.enqueue(v)
		doublylinkedcircularlist.enqueue(v)
	}

	singlylinkedlist.print()
	doublylinkedtailedlist.print()
	doublylinkedcircularlist.print()

	for doublylinkedcircularlist.len > 0 {
		fmt.Printf(
			"Dequeueing [%02d] from singly, doubly tailed, doubly circular: [%02d] [%02d] [%02d]\n",
			singlylinkedlist.peek().value,
			singlylinkedlist.dequeue().value,
			doublylinkedtailedlist.dequeue().value,
			doublylinkedcircularlist.dequeue().value)
	}
}

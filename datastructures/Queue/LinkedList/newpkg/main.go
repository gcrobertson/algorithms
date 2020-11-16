package main

import "fmt"

type element struct {
	next, prev *element
	list       *list
	value      interface{}
}

// next returns next element or nil
func (e *element) Next() *element {
	tp := e.next
	if tp == nil {
		return nil
	}
	if tp != e.list.root {
		return tp
	}
	return tp
}

func (e element) Previous() *element {
	if tp := e.prev; tp != nil && tp != e.list.root {
		return tp
	}
	return nil
}

// circularly linked list
// doubly linked linked list
func (l *list) Init() *list {
	l.root.prev = l.root
	l.root.next = l.root
	l.len = 0
	return l
}

type list struct {
	len  int
	root *element
}

// returns last element or nil
func (l *list) Back() *element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *list) insert(e, at *element) *element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = l
	l.len++
	return e
}

func (l *list) remove(e *element) *element {

	e.prev.next = e.next
	e.next.prev = e.prev

	// avoid memory leaks
	e.next = nil
	e.prev = nil
	e.list = nil

	l.len--
	return e
}

func (n *node) stack(val int) *node {
	t := &node{
		value: val,
		next:  n.next,
	}
	n.next = t

	return t
}

/* working singly linked list with no tail and:
 *
 * data structure: 	QUEUE
 *	enqueue: O(n)	because no end of linked list is calculable [add to end of a queue]
 *  dequeue: O(1)	remove first node swaps pointers in constant time
 *
 */
type singlylinkedlist struct {
	root *node
}

// singly linked list
type node struct {
	value int
	next  *node
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
		p := i.next
		t.next = p
		i.next = t
	}
	return t
}

func (l *singlylinkedlist) dequeue() *node {
	t := l.root
	l.root = t.next
	return t
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

func main() {

	fmt.Println("hello world, this is linked list learning!")

	singlylinkedlist := singlylinkedlist{}

	// va := new(list)
	// fmt.Printf("%T", singlylinkedlist)

	singlylinkedlist.enqueue(5)
	singlylinkedlist.enqueue(15)
	singlylinkedlist.enqueue(115)

	singlylinkedlist.print()

	// for link := singlylinkedlist.root; link != nil; link = link.next {

	// 	fmt.Printf("dequeue first in first out: %d\n", link.value)
	// }

	// for singlylinkedlist.root != nil {
	// 	fmt.Printf("dequeueing root: %d", singlylinkedlist.root.value)
	// 	r := singlylinkedlist.dequeue()
	// 	fmt.Printf("\tremoved %d\n", r.value)
	// }
}

func (l *list) enqueue(val int) *element {

	e := &element{
		value: val,
	}

	p := l.root.prev
	p.next = e

	n := l.root.next
	n.prev = e

	e.next = n
	e.prev = p

	l.len++

	return e
}

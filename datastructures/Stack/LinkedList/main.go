package main

import (
	"fmt"
)

/*
 *	Advantage of using a `Linked List` implementation of a `stack` over a `Slice` implementation:
 *	-	The stack can shrink or grow as much as need
 *	-	Array has restriction of maximum capacity which can lead to stack overflow
 */

// Node ... Singly Linked List
type Node struct {
	Val  int
	Link *Node
}

// Stack ...
type Stack struct {
	top *Node // top pointer which is the `head` of the stack.
}

// receiver needs a pointer because struct passed by copy
func (s *Stack) push(value int) {
	t := &Node{Val: value}
	if s.top == nil {
		s.top = t
	} else {
		t.Link = s.top
		s.top = t
	}
}

func (s Stack) isEmpty() bool {
	return s.top == nil
}

func (s Stack) peek() *Node {
	if s.top == nil {
		return nil
	}
	return s.top
}

// receiver needs a pointer because struct passed by copy
func (s *Stack) pop() *Node {
	if s.top == nil {
		return nil
	}
	t := s.top
	s.top = s.top.Link
	t.Link = nil
	return t
}

func (s Stack) display() {
	temp := s.top
	fmt.Printf("stack: ")
	for temp != nil {
		fmt.Printf("%v ", temp.Val)
		temp = temp.Link
	}
}

func main() {
	var s = Stack{}
	s.push(11)
	s.push(22)
	s.push(33)
	s.display()
	fmt.Printf("\ntop element is %d\n", s.top.Val)
	s.pop()
	s.pop()
	s.display()
	fmt.Printf("\ntop element is %d\n", s.top.Val)
}

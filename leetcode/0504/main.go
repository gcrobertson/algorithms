package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("convertToBase7: 15, Expect 21: 		Result: %v\n", convertToBase7(15))
	fmt.Printf("convertToBase7: 100, Expect 202: 	Result: %v\n", convertToBase7(100))
	fmt.Printf("convertToBase7: -7, Expect -10: 	Result: %v\n", convertToBase7(-7))
}

/*
Given an integer, return its base 7 string representation.
Example 1:
Input: 100
Output: "202"
Example 2:
Input: -7
Output: "-10"
Note: The input will be in range of [-1e7, 1e7].
*/

type node struct {
	val  int
	next *node
}

type stack struct {
	top *node
}

func (s *stack) push(v int) {
	t := &node{
		val: v,
	}
	if s.top == nil {
		s.top = t
	} else {
		t.next = s.top
		s.top = t
	}
}

func (s *stack) pop() *node {
	t := s.top
	s.top = s.top.next
	t.next = nil
	return t
}

func (s *stack) print() {
	t := s.top
	for t != nil {
		fmt.Printf("%v\t", t.val)
		t = t.next
	}
}

func (s *stack) peek() *node {
	return s.top
}

func convertToBase7(num int) string {

	if num == 0 {
		return "0"
	}

	var signed bool
	if num < 0 {
		signed = true
		num *= -1
	}

	stack := stack{}

	for num > 0 {
		stack.push(num % 7)
		num = num / 7
	}

	var r string
	if signed {
		r = "-" + r
	}

	for stack.peek() != nil {
		r = r + strconv.Itoa(stack.pop().val)
	}

	return r
}

// https://yourbasic.org/golang/build-append-concatenate-strings-efficiently/

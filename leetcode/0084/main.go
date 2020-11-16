package main

import "fmt"

func main() {

}

type node struct {
	value int
	next  *node
}

type stack struct {
	top *node
}

func (s *stack) peek() *node {
	return s.top
}

func (s *stack) push(value int) {
	n := &node{
		value: value,
	}
	if s.peek() == nil {
		s.top = n
	} else {
		n.next = s.top
		s.top = n
	}
}

func (s *stack) pop() *node {
	t := s.top
	s.top = s.top.next
	t.next = nil
	return t
}

func (s *stack) print() {
	if s.top == nil {
		return
	}
	t := s.top
	for t != nil {
		fmt.Printf("%v\t", t.value)
		t = t.next
	}
}

func largestRectangleArea(histogram []int) int {

	var index, height stack
	var area int

	input := histogram

	for i := 0; i < len(input); i++ {

		v := input[i]

		if index.peek() == nil || v > height.peek().value {

			index.push(i)
			height.push(input[i])

		} else if v < height.peek().value {

			var idx int

			for height.top != nil && v < height.peek().value {

				h := height.pop().value
				idx = index.pop().value
				w := i - idx

				size := h * w
				if size > area {
					area = size
				}
			}
			height.push(v)
			index.push(idx)
		}
	}

	for index.peek() != nil {
		h := height.pop().value
		w := len(input) - index.pop().value
		size := h * w
		if size > area {
			area = size
		}
	}

	return area
}

package main

import (
	"fmt"
	"math"
)

type node struct {
	value int
	left  *node
	right *node
}

type tree struct {
	tree   *node
	height int
}

type stack struct {
	top   *stack
	value *node
}

func main() {

	// preorder enqueue:
	/*
		t := tree{tree: &node{value: 1}}
		t.tree.left = &node{value: 2}
		t.tree.left.left = &node{value: 4}
		t.tree.left.right = &node{value: 5}
		t.tree.right = &node{value: 3}
	*/
	t := tree{tree: &node{value: 10}}
	t.tree.left = &node{value: 20}
	t.tree.left.left = &node{value: 30}
	t.tree.left.left.left = &node{value: 40}
	t.tree.left.right = &node{value: 50}
	t.tree.right = &node{value: 60}
	t.tree.right.right = &node{value: 70}
	t.tree.right.right.left = &node{value: 80}
	t.tree.right.right.right = &node{value: 90}
	t.tree.right.right.right.right = &node{value: 100}
	l := t.tree.leftist()

	fmt.Printf("leftist node: %v\n", l.value)

	s := &stack{}

	s.push(t.tree)
	s.push(t.tree.left)
	s.push(t.tree.left.left)
	// s.print()
	// s.pop()
	// s.pop()
	fmt.Printf("printing stack: ")
	s.print()
	fmt.Printf("\nprinting tree preorder:\n")
	printfPreOrder(t.tree)
	fmt.Printf("\nprinting tree inorder:\n")
	printfInOrder(t.tree)
	fmt.Printf("\nprinting tree postorder:\n")
	printfInOrder(t.tree)
	fmt.Printf("\n\ndepth of tree preorder: %v\n\n", depthfPreOrder(t.tree))
	fmt.Printf("\n\ndepth of tree inorder: %v\n\n", depthInOrder(t.tree))
	fmt.Printf("\n\ndepth of tree postorder: %v\n\n", depthPostOrder(t.tree))

}

// func stackPreOrder(n *node) {
// 	if n == nil {
// 		return
// 	}
// 	s := &stack{}

// 	t := n
// 	s.push(n)
// 	for t != nil {
// 		t = t.next
// 		s.push(n)
// 	}
// }

func depthfPreOrder(n *node) int {
	if n == nil {
		return 0
	}
	return int(math.Max(float64(depthfPreOrder(n.left)+1), float64(depthfPreOrder(n.right))+1))
}

func depthInOrder(n *node) int {
	if n == nil {
		return 0
	}
	left := depthInOrder(n.left)
	right := depthInOrder(n.right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func depthPostOrder(n *node) int {
	if n == nil {
		return 0
	}
	return int(math.Max(float64(depthPostOrder(n.right)), float64(depthPostOrder(n.left))) + 1)
}

func printfPreOrder(n *node) {
	if n == nil {
		return
	}
	fmt.Printf("%02d\t", n.value)
	printfPreOrder(n.left)
	printfPreOrder(n.right)
	return
}

func printfInOrder(n *node) {
	if n == nil {
		return
	}
	printfInOrder(n.left)
	fmt.Printf("%02d\t", n.value)
	printfInOrder(n.right)
	return
}

func printfPostOrder(n *node) {
	printfPostOrder(n.left)
	printfPostOrder(n.right)
	fmt.Printf("%02d\t", n.value)
}

func (s *stack) peek() int {
	if s.top == nil {
		return 0
	}
	return s.top.value.value
}

func (n *node) leftist() *node {
	l := n
	for l.left != nil {
		l = l.left
	}
	return l
}

func (s *stack) print() {
	t := s.top
	for t != nil {
		fmt.Printf("%v\t", t.value.value)
		t = t.top
	}
	fmt.Println()
}

func (s *stack) push(n *node) {
	t := &stack{value: n}
	if s.top == nil {
		s.top = t
	} else {
		t.top = s.top
		s.top = t
	}
}

func (s *stack) pop() int {
	t := s.top
	if t == nil {
		return 0
	}
	s.top = s.top.top
	return t.value.value
}

func (s *stack) isEmpty() bool {
	return s.top == nil
}

/*
func (n *node) heightDFS() int {
	if n == nil {
		return 0
	}
	var top *node
	var down int
	leftist := n
	for leftist != nil {
		top = leftist
		leftist = n.left
		down++
	}
	var height int
	var up int
	for leftist != nil {
		up, down = inorderDFS(leftist, 0, down)
		height = up - down
	}
	return height
}
func inorderDFS(top *node, up int, down int) (int, int) {
	if top == nil {
		return up, down
	}
	fmt.Printf("[inorder]\t%v", top.value)
	if top.left == nil && top.right == nil {
		up++
		return up, down
	}
	down++
	if top.left != nil {
		up, down := inorderDFS(top.left, up, down)
	}
	if top.right != nil {
		up, down := inorderDFS(top.left, up, down)
	}
	return up, down
}
*/

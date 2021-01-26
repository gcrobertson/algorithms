package main

import (
	"fmt"
	"reflect"
)

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
}

//			[1]
//			   \
//				[2]
//			   /
//			[3]
func ex1() {

	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	expect := []int{1, 2, 3}
	result := preorderTraversal(root)
	fmt.Printf("ex1(): Result [%v] matches expectation [%v]? [%v]\n", result, expect, reflect.DeepEqual(result, expect))
}

func ex2() {
	var root *TreeNode
	expect := []int{}
	result := preorderTraversal(root)
	fmt.Printf("ex2(): Result [%v] matches expectation [%v]? [%v]\n", result, expect, reflect.DeepEqual(result, expect))
}

func ex3() {

	root := &TreeNode{Val: 1}

	expect := []int{1}
	result := preorderTraversal(root)
	fmt.Printf("ex3(): Result [%v] matches expectation [%v]? [%v]\n", result, expect, reflect.DeepEqual(result, expect))
}

func ex4() {

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}

	expect := []int{1, 2}
	result := preorderTraversal(root)
	fmt.Printf("ex4(): Result [%v] matches expectation [%v]? [%v]\n", result, expect, reflect.DeepEqual(result, expect))
}

// TreeNode ... defined by Leetcode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*	DFS Solution #1
 *
 *
 *
 *	Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
 *  Memory Usage: 2.1 MB, less than 39.92% of Go online submissions for Binary Tree Preorder Traversal.
 */
func preorderTraversalDFS1(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	var result []int
	result = append(result, root.Val)

	if root.Left != nil {
		result = append(result, preorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		result = append(result, preorderTraversal(root.Right)...)
	}

	return result
}

/*	DFS Solution #2
 *
 *
 *
 *	Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
 *  Memory Usage: 2.1 MB, less than 39.92% of Go online submissions for Binary Tree Preorder Traversal.
 */
func preorderTraversalDFS2(root *TreeNode) []int {
	// O(1)
	var result []int
	// O(n)?
	preorder(root, &result)
	// O(1)
	return result
}

func preorder(root *TreeNode, r *[]int) {
	// O(1)
	if root != nil {
		// O(1)
		*r = append(*r, root.Val)

		preorder(root.Left, r)
		preorder(root.Right, r)
	}
}

/*	BFS Solution #1
 *
 *
 *
 *	Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
 *  Memory Usage: 2 MB, less than 39.84% of Go online submissions for Binary Tree Preorder Traversal.
 */
func preorderTraversal(root *TreeNode) []int {

	// O(1)
	if root == nil {
		return []int{}
	}

	var result []int

	stack := treeStack{}
	stack.add(root)

	// O(n)?
	for stack.isEmpty() == false {

		node := stack.remove()

		result = append(result, node.Val.Val)

		if node.Val.Right != nil {
			stack.add(node.Val.Right)
		}

		if node.Val.Left != nil {
			stack.add(node.Val.Left)
		}
	}

	// O(1)
	return result
}

// Stack element for stack data structure `treeStack`
type treeStackNode struct {
	Val  *TreeNode
	Link *treeStackNode
}

// stack data structure for `TreeNode` element
type treeStack struct {
	top *treeStackNode
}

// push onto the stack
func (ts *treeStack) add(val *TreeNode) {

	node := &treeStackNode{
		Val: val,
	}

	if ts.top == nil {
		ts.top = node
	} else {
		node.Link = ts.top
		ts.top = node
	}
}

// pop from the stack
func (ts *treeStack) remove() *treeStackNode {
	if ts.top == nil {
		return nil
	}

	t := ts.top
	ts.top = ts.top.Link
	t.Link = nil

	return t
}

// isEmpty
func (ts *treeStack) isEmpty() bool {
	return ts.top == nil
}

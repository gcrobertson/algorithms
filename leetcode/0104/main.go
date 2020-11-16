package main

import "fmt"

func main() {

}

// TreeNode ... given as a data structure by Leetcode
type TreeNode struct {
	val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	// base case
	if root == nil {
		return 0
	}
	// create an empty queue for `level order traversal`
	//  - to enqueue use built-in append()
	//  - to dequeue slice off the first element
	queue := []*TreeNode{}

	// enqueue root and initialize height
	queue = append(queue, root)
	var height int

	for {

		// nodeCount [queue size] indicates the number of nodes at the current level.
		var nodeCount = len(queue)

		fmt.Printf("Outside for, the nodeCount is %v\n", nodeCount)

		if nodeCount == 0 {
			return height
		}
		height++

		// dequeue all nodes of current level
		// and enqueue all nodes of next level
		for nodeCount > 0 {
			node := queue[0]
			queue = queue[1:] // dequeue

			//fmt.Printf("first queue should be empty: %v\n", queue)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			nodeCount--

			fmt.Printf("Inside for, nodeCount is %v and queue length should be 3 at some point: %v\n", nodeCount, len(queue))
		}
	}
}

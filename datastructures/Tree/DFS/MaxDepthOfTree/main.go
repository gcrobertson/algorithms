package main

import "fmt"

// https://www.teamblind.com/post/New-Year-Gift---Curated-List-of-Top-75-LeetCode-Questions-to-Save-Your-Time-OaM1orEU

// https://leetcode.com/problems/maximum-depth-of-binary-tree/

//TreeNode ... recursive data structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
// Question for Prakriti: The maximum depth is the same as the height of the binary tree?
func maxDepth(root *TreeNode) int {

	// Recursively calculate the height of the tree to the left of the root.
	// Recursively calculate the height of the tree to the right of the root.
	// Pick the larger height from the two answers and add one to it (to account for the root node).

	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// Insert ... pointer receiver
func (n *TreeNode) Insert(val int) {
	if n.Val > val {
		if n.Left == nil {
			n.Left = &TreeNode{Val: val}
		} else {
			n.Left.Insert(val)
		}
	} else if val > n.Val {
		if n.Right == nil {
			n.Right = &TreeNode{Val: val}
		} else {
			n.Right.Insert(val)
		}
	}
}

func main() {

	tree := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	tree.Insert(9)
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(7)

	fmt.Printf("The maximum depth of the tree is: %v\n", maxDepth(tree))

}

// original problem statement:
// maximum depth of a binary tree:
// - exercise with prakriti here: 	https://leetcode.com/problems/maximum-depth-of-binary-tree/
// - go solution here:

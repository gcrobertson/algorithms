package main

import (
	"fmt"
)

func main() {

	t := &TreeNode{Val: 1}
	t.Left = &TreeNode{Val: 2}
	t.Right = &TreeNode{Val: 3}
	t.Left.Left = &TreeNode{Val: 4}
	t.Left.Right = &TreeNode{Val: 5}

	result := diameterOfBinaryTree(t)

	fmt.Printf("Result: [%v]\n", result)
}

// TreeNode ...
type TreeNode struct {
	Val   int
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

// Runtime: 4 ms, faster than 95.24% of Go online submissions for Diameter of Binary Tree.
// Memory Usage: 4.6 MB, less than 59.31% of Go online submissions for Diameter of Binary Tree.
/*
var ans int


func diameterOfBinaryTree(root *TreeNode) int {
	ans = 1
	depth(root)
	return ans - 1
}

func depth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	l := depth(root.Left)
	r := depth(root.Right)
	ans = int(math.Max(float64(ans), float64(l)+float64(r)+1))
	return int(math.Max(float64(l), float64(r)) + 1)
}
*/

// Runtime: 4 ms, faster than 95.24% of Go online submissions for Diameter of Binary Tree.
// Memory Usage: 4.6 MB, less than 59.31% of Go online submissions for Diameter of Binary Tree.
/*
var ans int

func diameterOfBinaryTree(root *TreeNode) int {

	ans = 0

	recursion(root)

	return ans
}

func recursion(root *TreeNode) int {

	//
	//
	if root == nil {
		return 0
	}

	//
	//
	l := recursion(root.Left)
	r := recursion(root.Right)

	//
	//
	ans = int(math.Max(float64(ans), float64(l)+float64(r)))

	ret := int(math.Max(float64(l), float64(r)) + 1)

	// fmt.Printf("val: %v\tl: %v\tr: %v\tans: %v\tret: %v\n", root.Val, l, r, ans, ret)

	//
	//
	return ret
}
*/

// solution courtesy: n960321
func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	_, count := nextLevel(root, &result)
	if count > result {
		result = count
	}
	return result
}

func nextLevel(node *TreeNode, result *int) (d, w int) {
	if node == nil {
		return 0, 0
	}
	rLevel, rCount := nextLevel(node.Left, result)
	lLevel, lCount := nextLevel(node.Right, result)

	switch {
	case rCount > *result:
		*result = rCount
	case lCount > *result:
		*result = rCount
	}

	if rLevel > lLevel {
		return rLevel + 1, rLevel + lLevel
	}
	return lLevel + 1, rLevel + lLevel

}

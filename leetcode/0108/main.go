package main

import (
	"fmt"
	"math"
)

//	https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

//TreeNode ...
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
func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	midIndex := int(math.Floor(float64(len(nums) / 2)))
	mid := nums[midIndex]

	node := &TreeNode{Val: mid}
	node.Left = sortedArrayToBST(nums[:midIndex])
	node.Right = sortedArrayToBST(nums[midIndex+1:])

	return node
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}

	tree := sortedArrayToBST(nums)

	fmt.Println("inorder:")
	inOrder(tree)
	fmt.Println("preorder:")
	preOrder(tree)
	fmt.Println("postorder:")
	postOrder(tree)

	fmt.Println()
}

func preOrder(t *TreeNode) {
	if t == nil {
		return
	}
	fmt.Printf("%v\t", t.Val)
	preOrder(t.Left)
	preOrder(t.Right)
	return
}

func inOrder(t *TreeNode) {
	if t == nil {
		return
	}
	inOrder(t.Left)
	fmt.Printf("%v\t", t.Val)
	inOrder(t.Right)
	return
}

func postOrder(t *TreeNode) {
	if t == nil {
		return
	}
	postOrder(t.Left)
	postOrder(t.Right)
	fmt.Printf("%v\t", t.Val)
	return
}

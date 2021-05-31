package main

import (
	"fmt"
	"reflect"
)

func main() {

	t1 := &TreeNode{Val: 1}
	t1.Left = &TreeNode{Val: 3}
	t1.Right = &TreeNode{Val: 2}
	t1.Left.Left = &TreeNode{Val: 5}

	t2 := &TreeNode{Val: 2}
	t2.Left = &TreeNode{Val: 1}
	t2.Right = &TreeNode{Val: 3}
	t2.Left.Right = &TreeNode{Val: 4}
	t2.Right.Right = &TreeNode{Val: 7}

	result := mergeTrees(t1, t2)

	merged := &TreeNode{Val: 3}
	merged.Left = &TreeNode{Val: 4}
	merged.Left.Left = &TreeNode{Val: 5}
	merged.Left.Right = &TreeNode{Val: 4}
	merged.Right = &TreeNode{Val: 5}
	merged.Right.Right = &TreeNode{Val: 7}

	fmt.Printf("Example 1? %v\n", reflect.DeepEqual(merged, result))

	printPreOrder(result)
	fmt.Println()
	printPreOrder(merged)
	fmt.Println()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	root := &TreeNode{Val: root1.Val + root2.Val}
	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)

	return root
}

func printPreOrder(root *TreeNode) {

	if root == nil {
		return
	}

	fmt.Printf("%v\t", root.Val)

	printPreOrder(root.Left)
	printPreOrder(root.Right)
}

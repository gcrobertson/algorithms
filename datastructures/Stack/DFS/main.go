package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/maximum-depth-of-binary-tree/discuss/34195/Two-Java-Iterative-solution-DFS-and-BFS

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newNode(data int) *TreeNode {
	return &TreeNode{
		Val:   data,
		Left:  nil,
		Right: nil,
	}
}

func main() {

	root := newNode(1)
	root.Left = newNode(2)
	root.Right = newNode(3)
	root.Left.Left = newNode(4)
	root.Left.Right = newNode(5)

	fmt.Printf("Height of tree: %v\n", treeHeight(root))

}

func push(stack []*TreeNode, node *TreeNode) {
	stack = append(stack, node)
}

func pop(stack []*TreeNode) *TreeNode {

	node := stack[len(stack)-1] // pop
	stack = stack[:len(stack)-1]

	return node
}

func treeHeight(root *TreeNode) int {

	if root == nil {
		return 0
	}

	stack := []*TreeNode{}
	value := []int{}

	push(stack, root)        // push
	value = append(value, 1) // push

	var max int

	for len(stack) > 0 {

		node := pop(stack)
		temp := value[len(value)-1] // pop
		value = value[:len(value)-1]

		max = int(math.Max(float64(max), float64(temp)))

		if node.Left != nil {
			push(stack, node.Left)        // push
			value = append(value, temp+1) // push
		}

		if node.Right != nil {
			push(stack, node.Right)       // push
			value = append(value, temp+1) // push
		}
	}

	return max
}

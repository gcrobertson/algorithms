package main

import "fmt"

func main() {

	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 9}

	root.print()

	inverted := invertTree(root)

	inverted.print()
}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) print() {

	q := []*TreeNode{n}

	for len(q) > 0 {

		dq := q[0]

		q = q[1:]

		fmt.Printf("%v\t", dq.Val)

		if dq.Left != nil {
			q = append(q, dq.Left)
		}
		if dq.Right != nil {
			q = append(q, dq.Right)
		}
	}

	fmt.Println()
}

/*
Input:

     4
   /   \
  2     7
 / \   / \
1   3 6   9

Output:

     4
   /   \
  7     2
 / \   / \
9   6 3   1

*/
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Invert Binary Tree.
// Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Invert Binary Tree.
func invertTreeBFS(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		dq := q[0]
		q = q[1:]

		l := dq.Left
		r := dq.Right

		dq.Left = r
		dq.Right = l

		if dq.Left != nil {
			q = append(q, dq.Left)
		}
		if dq.Right != nil {
			q = append(q, dq.Right)
		}
	}

	return root
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Invert Binary Tree.
// Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Invert Binary Tree.
func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

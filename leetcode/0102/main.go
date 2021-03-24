package main

func main() {

}

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
func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{}
	q = append(q, root)
	var r [][]int // result

	for len(q) > 0 { // as long as there is a queue
		t := []int{} // all values of a tree row

		for l := len(q); l > 0; l-- {
			dq := q[0]
			q = q[1:]

			t = append(t, dq.Val) // add to tree row
			if dq.Left != nil {
				q = append(q, dq.Left) // add to queue
			}
			if dq.Right != nil {
				q = append(q, dq.Right) // add to queue
			}
		}
		r = append(r, t)
	}

	return r
}

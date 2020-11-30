package main

func main() {

}

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

/*

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	lc := lowestCommonAncestor(root.Left, p, q)
	rc := lowestCommonAncestor(root.Right, p, q)

	if lc == nil {
		return rc
	}
	if rc == nil {
		return lc
	}
	return root
}

*/

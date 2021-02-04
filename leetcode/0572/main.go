package main

import (
	"fmt"
)

func main() {
	ex1()
	// ex2()
}

func ex1() {

	tree := &TreeNode{Val: 3}
	tree.Left = &TreeNode{Val: 4}
	tree.Right = &TreeNode{Val: 5}
	tree.Left.Left = &TreeNode{Val: 1}
	tree.Left.Right = &TreeNode{Val: 2}

	subtree := &TreeNode{Val: 4}
	subtree.Left = &TreeNode{Val: 1}
	subtree.Right = &TreeNode{Val: 2}

	expect := true

	// For some strange reason, it works when the subtree & tree are inverted ... WHY?
	result := isSubtree(tree, subtree)

	fmt.Printf("Is the result [%v] correct? [%v]\n", result, result == expect)
}

func ex2() {

	tree := &TreeNode{Val: 3}
	tree.Left = &TreeNode{Val: 4}
	tree.Right = &TreeNode{Val: 5}
	tree.Left.Left = &TreeNode{Val: 1}
	tree.Left.Right = &TreeNode{Val: 2}
	tree.Left.Right.Left = &TreeNode{Val: 0}

	subtree := &TreeNode{Val: 4}
	subtree.Left = &TreeNode{Val: 1}
	subtree.Right = &TreeNode{Val: 2}

	expect := false
	result := isSubtree(subtree, tree)

	fmt.Printf("Is the result [%v] expected? [%v]\n", result, result == expect)
}

// TreeNode defined by Leetcode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {

	if s == nil || t == nil {
		return s == t
	}

	// fmt.Printf("s.Val: %v\tt.Val: %v\n", s.Val, t.Val)

	if s.Val == t.Val && isIdentical(s, t) {
		return true
	}

	/*
	 *
	 *
	 *
	 */
	// fmt.Printf("s.Left: [%v]\t t: [%v]\n", s.Left, t)

	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

// Remember, preorder = Root,Left,Right
func isIdentical(s, t *TreeNode) bool {
	if s == nil || t == nil {
		return s == t
	}
	return s.Val == t.Val && isIdentical(s.Left, t.Left) && isIdentical(s.Right, t.Right)
}

/*
func isSubtree(s *TreeNode, t *TreeNode) bool {
	tree1 := preorder(s, true)
	tree2 := preorder(t, true)
	return strings.Contains(tree2, tree1)
}

func preorder(t *TreeNode, left bool) string {
	if t == nil {
		if left == true {
			return "lnull"
		}
		return "rnull"
	}

	return "#" + strconv.Itoa(t.Val) + " " + preorder(t.Left, true) + " " + preorder(t.Right, false)
}
*/

/*
func isSubtree(s *TreeNode, t *TreeNode) bool {

	// so, we have a subtree... we need to find the `root` of the subtree inside of the `tree`.

	// once we find a `match` for the root of the subtree, traverse the subtree and make sure that
	// the tree at that stage has the same children as the subtree
	// until both the tree and subtree are exhausted ... sounds easy, sounds hard to implement lol

	// a queue seems to be the easiest way to do it ...

	q := []*TreeNode{t}

	for len(q) > 0 {

		dq := q[0]
		q = q[1:]

		if dq.Val == s.Val {
			l := recurse(s.Left, dq.Left)
			r := recurse(s.Right, dq.Right)

			if l == true && r == true {
				return true
			}
		}

		if dq.Left != nil {
			q = append(q, dq.Left)
		}
		if dq.Right != nil {
			q = append(q, dq.Right)
		}
	}

	return false
}

func recurse(s *TreeNode, t *TreeNode) bool {

	if (s.Left == nil && t.Left != nil) || (s.Right == nil && t.Right != nil) {
		return false
	}
	if s.Left == nil && t.Left == nil && s.Right == nil && t.Right == nil {
		return true
	}

	if s.Val != t.Val {
		return false
	}

	l := recurse(s.Left, t.Left)
	r := recurse(s.Right, t.Right)

	return l == true && r == true
}
*/

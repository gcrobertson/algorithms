package main

import (
	"fmt"
)

func main() {

	// [1]
	// [1,2,2,1]
	// [1,2,6,3,4,5,6]

	head := &ListNode{
		Val: 1,
	}

	h := head

	// for _, v := range []int{2, 6, 3, 4, 5, 6} {
	// for _, v := range []int{1, 2, 2, 1} {
	for _, v := range []int{2, 6, 3, 4, 5, 6} {

		t := &ListNode{
			Val: v,
		}

		h.Next = t
		h = h.Next
	}

	fmt.Println("Before")
	outputElements(head)

	result := removeElements(head, 6)
	fmt.Println("After")
	outputElements(result)
}

func outputElements(head *ListNode) {

	if head == nil {
		return
	}

	r := head
	fmt.Printf("ListNode: [%v]\n", r.Val)
	for r.Next != nil {
		r = r.Next
		fmt.Printf("ListNode: [%v]\n", r.Val)
	}
}

// ListNode defined by Leetcode
type ListNode struct {
	Val  int
	Next *ListNode
}

// Runtime: 4 ms, faster than 99.54% of Go online submissions for Remove Linked List Elements.
// Memory Usage: 4.7 MB, less than 100.00% of Go online submissions for Remove Linked List Elements.
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	cur := dummy

	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}

// Runtime: 8 ms, faster than 87.16% of Go online submissions for Remove Linked List Elements.
// Memory Usage: 4.7 MB, less than 100.00% of Go online submissions for Remove Linked List Elements.
func removeElementsSolution1(head *ListNode, val int) *ListNode {

	if head == nil {
		return nil
	}

	p := head // current node

	// handle first node
	for p != nil && p.Val == val {
		head = head.Next
		p = head
	}

	// handle all other nodes
	for p != nil {
		if p.Next == nil {
			break
		}
		for p.Next != nil && p.Next.Val == val {
			p.Next = p.Next.Next
		}
		p = p.Next
	}
	return head
}

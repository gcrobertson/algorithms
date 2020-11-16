package main

import (
	"fmt"
)

// Merge two sorted linked lists and return it as a new sorted list.
// The new list should be made by splicing together the nodes of the first two lists.

// Example: Input: 	1>2>4, 1>3>4
// 			Output:	1>1>2>3>4>4

func main() {
	var list1 = &ListNode{
		Val: 1,
	}
	list1.push(2)
	list1.push(4)
	list1.print()

	var list2 = &ListNode{
		Val: 1,
	}
	list2.push(3)
	list2.push(4)
	list2.print()

	result := mergeTwoLists(list1, list2)
	result.print()
}

// is the linked list sorted or not?
// binary search algorithm to cut it into log(n) time
// dummy head approach = keeping a reference to a list
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2
			l2 = nil
			break
		}
		if l2 == nil {
			cur.Next = l1
			l1 = nil
			break
		}
		if l1.Val < l2.Val {
			cur.Next = l1
			cur, l1 = cur.Next, l1.Next
		} else {
			cur.Next = l2
			cur, l2 = cur.Next, l2.Next
		}
	}
	return res.Next
}

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) push(n int) {
	t := l
	for t.Next != nil {
		t = t.Next
	}
	t.Next = &ListNode{
		Val: n,
	}
}

func (l *ListNode) print() {
	for t := l; t != nil; t = t.Next {
		fmt.Printf("%v\t", t.Val)
	}
	fmt.Println()
}

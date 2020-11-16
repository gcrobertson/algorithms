package main

import "fmt"

func main() {

	var input = ListNode{Val: 1}
	input.push(2)
	input.push(3)
	input.push(4)
	input.push(5)

	result := oddEvenList(&input)
	result.print()

}

func (l *ListNode) print() {
	t := l
	for t != nil {
		fmt.Printf("%v\t", t.Val)
		t = t.Next
	}
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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	odd := head
	even := odd.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
}

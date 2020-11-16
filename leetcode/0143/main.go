package main

func main() {}

// ListNode declared by Leetcode
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Find the middle
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse the right part
	head2 := slow.Next
	slow.Next = nil
	var prev *ListNode
	for head2 != nil {
		temp := head2.Next
		head2.Next = prev
		prev = head2
		head2 = temp
	}

	cur1, cur2 := head, prev
	for cur2 != nil {
		next1, next2 := cur1.Next, cur2.Next
		cur1.Next, cur2.Next = cur2, next1
		cur1, cur2 = next1, next2
	}
	return
}

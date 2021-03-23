/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	seen := make(map[*ListNode]bool)

	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = true
		head = head.Next
	}
	return false
}

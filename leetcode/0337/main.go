package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	// [3,2,3,null,3,null,1]

	tree := &TreeNode{Val: 3}
	tree.Left = &TreeNode{Val: 2}
	tree.Right = &TreeNode{Val: 3}
	tree.Left.Left = nil
	tree.Left.Right = &TreeNode{Val: 3}
	tree.Right.Left = nil
	tree.Right.Right = &TreeNode{Val: 1}

	result := rob(tree)

	fmt.Printf("the result is %v\n", result)
}

// TreeNode defined in Leetcode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(tree *TreeNode) int {
	if tree == nil {
		return 0
	}
	if tree.Left == nil && tree.Right == nil {
		return tree.Val
	}

	queue := []*TreeNode{tree}

	var prevMax, currMax int

	for {

		fmt.Printf("prevMax: %v\tcurrMax:%v\t\n", prevMax, currMax)

		var currLevel, temp int

		length := len(queue)

		fmt.Printf("length of queue:%v\n", len(queue))

		for length > 0 {

			temp = currMax

			length--
			dequeue := queue[0]
			queue := queue[1:]

			fmt.Printf("length of queue:%v\n", len(queue))

			fmt.Printf("dequeue: %v\t\n", dequeue)

			if dequeue.Left != nil {
				queue = append(queue, dequeue.Left)
				fmt.Println("appended dequeue.Left!")
			}
			if dequeue.Right != nil {
				queue = append(queue, dequeue.Right)
				fmt.Println("appended dequeue.Right!")
			}

			currLevel += dequeue.Val

			fmt.Printf("length of queue:%v\n", len(queue))
		}

		currMax = int(math.Max(float64(currLevel+prevMax), float64(currMax)))

		prevMax = temp

		if len(queue) == 0 {
			return currMax
		}

		if prevMax > 10 {
			os.Exit(5)
		}
	}
}

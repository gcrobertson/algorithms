package main

import "fmt"

/*
 *	Rotate Array
 *
 *	Given an array, rotate the array to the right by k steps, where k is non-negative.
 *
 *	Example 1:
 *		Input: 	nums = [1,2,3,4,5,6,7], k = 3
 *		Output: [5,6,7,1,2,3,4]
 */

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Printf("output: %v\n", nums)
}

// using a copy, O(n) + O(n)
func rotate(nums []int, k int) {
	result := make([]int, len(nums), len(nums)) // O(1)
	for i := 0; i < len(nums); i++ {            // O(n)
		/*	(i+k)%len(nums)	= key

			len(nums)	7

			k			3

			i		key mod
			0		3	3%7
			1		4	4%7
			2		5	5%7
			3		6	6%7
			4		0	7%7
			5		1	8%7
			6		2	9%7
			7		3  10%7
			8		4  11%7
			9		5  12%7
		*/
		result[(i+k)%len(nums)] = nums[i] // O(1)
	}
	copy(nums, result) // O(n)
}

// rotating in place, O(n*n)
func rotate2(nums []int, k int) {
	k = k % len(nums) // O(1)
	for k > 0 {       // O(n)
		n := nums[len(nums)-1]  // O(1) copy last element
		copy(nums[1:], nums[:]) // O(n) shift all elements forward 1 [this deletes the last element]
		nums[0] = n             // O(1) set last element value to first element
		k--                     // O(1) decrement the number of additional shifts to perform
	}
}

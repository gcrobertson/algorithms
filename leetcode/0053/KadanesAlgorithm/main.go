package main

import (
	"fmt"
	"math"
)

func main() {

	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expect := 6 // 4, -1, 2, 1
	result := kadanesAlgorithm(input)

	fmt.Printf("Does the result [%v] match the expectation [%v]?\t [%v]\n", result, expect, result == expect)
}

// Approach 3: Kadane's Algorithm
func kadanesAlgorithm(nums []int) int {

	n := len(nums)
	maxSum := nums[0]

	// fmt.Printf("Initial nums array: %v\n", nums)

	for i := 1; i < n; i++ {

		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}

		maxSum = int(math.Max(float64(nums[i]), float64(maxSum)))

		fmt.Printf("at i: %v maxSum: %v nums: %v\n", i, maxSum, nums)
	}

	// fmt.Printf("Modified nums array: %v\n", nums)

	return maxSum
}

// index 0:			 0  1   2  3   4  5  6   7  8
// array: 			-2, 1, -3, 4, -1, 2, 1, -5, 4
// max:				-2
//
// nums array only updated on index 2, 4, 5, 7.
//
// index 1:			 0  1   2  3   4  5  6   7  8
// array: 			-2, 1, -3, 4, -1, 2, 1, -5, 4
// max:				 1
//
// index 2:			 0  1   2  3   4  5  6   7  8
// array: 			-2, 1, -2*,4, -1, 2, 1, -5, 4
// max: 			 1
//
// index 3:			 0  1   2  3   4  5  6   7  8
// array: 			-2, 1, -2, 4, -1, 2, 1, -5, 4
// max: 			 4
//
// index 4:			 0  1   2  3  4   5  6   7  8
// array: 			-2, 1, -2 ,4, 3*, 2, 1, -5, 4
// max: 			 4
//
// index 5:			 0  1   2  3   4  5  6   7  8
// array: 			-2, 1, -2 ,4,  3, 5*,1, -5, 4
// max: 			 5
//
// index 6:			 0  1   2  3   4  5  6   7  8
// array: 			-2, 1, -2 ,4,  3, 5 ,6*,-5, 4
// max: 			 6
//
// index 7:			 0  1   2  3   4  5  6   7  8
// array: 			-2, 1, -2, 4,  3, 5, 6,  1*,4
// max: 6
//
// index 8:			 0  1   2  3   4  5  6   7  8
// array: 			-2, 1, -2, 4,  3, 5, 6,  1, 5*
// max: 6
//

// https://leetcode.com/problems/maximum-subarray/solution/

//  The problem to find sum or maximum or minimum in an entire array or in a fixed-size sliding window.

//  There are 2 standard DP approaches suitable for arrays:
//  - Constant space one. 	Move along the array and modify the array itself.
//  - Linear space one.		First move in the direction left to right, and then in the direction right to left. Combine the results. See: [https://leetcode.com/problems/sliding-window-maximum/solution/]

//

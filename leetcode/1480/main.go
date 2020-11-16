package main

import (
	"fmt"
)

// Given an array `nums`

// We define a running sum of an array as: runningSum[i] = sums(nums[0]...nums[i])
// Return the running sum of nums

// Example 1:
// Input: nums = [1,2,3,4]
// Output: [1,3,6,10]

// ➜  1480.RunningSumOf1dArray git:(master) ✗ go run main.go
// expectation: [1 3 6 10] result: [1 3 6 10]

func main() {

	example1 := []int{1, 2, 3, 4}
	output1 := []int{1, 3, 6, 10}
	result1 := runningSum(example1)

	fmt.Printf("expectation: %v result: %v\n", output1, result1)
}

// O(n)
// func runningSum(nums []int) []int {
// 	if len(nums) == 0 { // O(1)
// 		return nums // O(1)
// 	}
// 	result := make([]int, len(nums)) // O(1)
// 	var i int                        // O(1)
// 	result[i] = nums[i]              // O(1)
// 	for i := 1; i < len(nums); i++ { // O(n)
// 		result[i] = result[i-1] + nums[i] // O(1)
// 	}
// 	return result // O(1)
// }

func runningSum(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] + nums[i]
	}
	return result
}

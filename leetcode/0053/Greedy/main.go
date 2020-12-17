package main

import (
	"fmt"
	"math"
)

func main() {

	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expect := 6 // 4, -1, 2, 1
	result := maxSubArray(input)

	fmt.Printf("Does the result [%v] match the expectation [%v]?\t [%v]\n", result, expect, result == expect)
}

// Greedy approach can be solved in linear time.
// See: `Super Washing Machines`
// See: `Gas Problem`
func maxSubArray(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	curMax, totMax := nums[0], nums[0]
	for _, num := range nums {
		curMax = int(math.Max(float64(curMax+num), float64(num)))
		totMax = int(math.Max(float64(curMax), float64(totMax)))
	}

	return totMax

	//  index:		 0		 1		2		3		4		5		6		7		8
	//
	//	value:		-2	 	 1	   -3		4	   -1		2		1		-5		4
	//
	//	curSum:		-2		 1	   -2		4		3		5		6		1		5
	//	maxSum: 	-2		 1		1		4		4		5		6		6		6
}

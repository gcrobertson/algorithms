package main

import (
	"fmt"
	"math"
)

func main() {

	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expect := 6 // 4, -1, 2, 1
	result := divideAndConquer(input)

	fmt.Printf("Does the result [%v] match the expectation [%v]?\t [%v]\n", result, expect, result == expect)
}

// Approach 1: Divide & Conquer
// define the base problem
// split the problem into subproblems and solve them recursively
// merge the solutions for the subproblems to obtain the solution for the original problem

// See: https://leetcode.com/explore/learn/card/recursion-ii/470/divide-and-conquer/

// @TODO: Continue to explore, this is a very difficult algorithm to try to understand.
func divideAndConquer(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	return divideAndConquerHelper(nums, 0, len(nums)-1)
}

func divideAndConquerHelper(nums []int, left, right int) int {

	if left == right {
		return nums[left]
	}

	midpoint := (left + right) / 2

	// fmt.Printf("Helper.\t\tValues: nums [%v], left: [%v], right: [%v], p: [%v]\n", nums, left, right, midpoint)

	fmt.Printf("leftSum:\tleft: [%v]\t right: [%v]\n", left, midpoint)

	leftSum := divideAndConquerHelper(nums, left, midpoint)

	fmt.Printf("leftSum:\t [%v]\n", leftSum)
	fmt.Printf("rightSum:\tleft: [%v]\t right: [%v]\n", midpoint+1, right)

	rightSum := divideAndConquerHelper(nums, midpoint+1, right)

	fmt.Printf("rightSum:\t [%v]\n", rightSum)

	crossSum := divideAndConquerCrossSum(nums, left, right, midpoint)

	fmt.Printf("crossSum:\t [%v]\n", crossSum)

	// merge all subproblems solutions
	return int(math.Max(math.Max(float64(leftSum), float64(rightSum)), float64(crossSum)))
}

func divideAndConquerCrossSum(nums []int, left, right, p int) int {

	// fmt.Printf("CrossSum.\tValues: nums [%v], left: [%v], right: [%v], p: [%v]\n", nums, left, right, p)

	if left == right {
		return nums[left]
	}

	var leftSubSum, currSum int

	for i := p; i > left-1; i-- {
		currSum += nums[i]
		leftSubSum = int(math.Max(float64(leftSubSum), float64(currSum)))
	}

	var rightSubSum int
	currSum = 0

	for i := p + 1; i < right+1; i++ {
		currSum += nums[i]
		rightSubSum = int(math.Max(float64(rightSubSum), float64(currSum)))
	}

	return leftSubSum + rightSubSum
}

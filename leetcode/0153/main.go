package main

import "fmt"

func main() {

	// ex1()
	// ex2()
	// ex3()
	ex4()
}

func ex1() {
	input := []int{3, 4, 5, 1, 2}
	result := findMin(input)
	fmt.Printf("Does the result %v match the expectation of 1? [%v]\n", result, result == 1)
}

func ex2() {
	input := []int{4, 5, 6, 7, 0, 1, 2}
	result := findMin(input)
	fmt.Printf("Does the result %v match the expectation of 0? [%v]\n", result, result == 0)
}

func ex3() {
	input := []int{11, 13, 15, 17}
	result := findMin(input)
	fmt.Printf("Does the result %v match the expectation of 11? [%v]\n", result, result == 11)
}

func ex4() {
	input := []int{5, 1, 2, 3, 4}
	result := findMin(input)
	fmt.Printf("Does the result %v match the expectation of 1? [%v]\n", result, result == 1)
}

// Binary Search ...
// We find the mid-point of the array and decide to either search on the left or right

// Since the given array is sorted, we can make use of binary search.

// Because the array is rotated, simply applying binary search won't work here but we can apply a modified version of it.

// The inflection point here:
// When the last element is greater than the fist element, the array is in order
// When the first element is greater than the last element, the array is rotated

// Algorithm:
// 1. Find the mid element of the array
// 2. If mid element > first element, these elements are in order.  We need to search to the right to find the inflection point.
// 3. If mid element < first element, we know the inflection point is to the left side.

// We can stop the search when either of these conditions is met:
// 1. nums[mid] > nums[mid+1]
// 2. nums[mid-1] > nums[mid]

// O(log n), because it is dividing the search in half each time.
// TAG: `Binary Search`
func findMin(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	var left int
	right := len(nums) - 1

	if nums[left] < nums[right] {
		// array is already in ascending order
		return nums[left]
	}

	// Being binary search
	for right >= left {

		// Find the middle element
		mid := left + (right-left)/2

		// fmt.Printf("left: %v\t right: %v\t mid: %v\n", left, right, mid)

		if nums[mid] > nums[mid+1] { // mid+1 < len(nums) &&
			return nums[mid+1]
		}

		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		if nums[mid] > nums[0] {
			left = mid + 1 // smallest value is somewhere to the right of middle
		} else {
			right = mid - 1 // smallest value is somewhere to the left of middle
		}
	}

	return -1
}

// O(n), faster than brute force because it will exit early in the linear search.
func findMinON(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if nums[0] < nums[len(nums)-1] { // the array has been rotated to the point it is back in ascending order
		return nums[0]
	}

	// going forward ...
	min := nums[0]
	for i := 0; i < len(nums)-1; i++ {

		if nums[i] > nums[i+1] {
			min = nums[i+1]
			break
		}
	}

	return min

	/*

		// going backward ...
		min := nums[len(nums)-1]
		for i := len(nums) - 1; i > 0; i-- {
			if nums[i-1] > nums[i] {
				min = nums[i]
				break
			}
		}
		return min

	*/

	/*
		// brute force

		min := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] < min {
				min = nums[i]
			}
		}

		return min

	*/
}

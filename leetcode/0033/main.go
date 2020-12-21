package main

import "fmt"

func main() {

	ex1()
	ex2()
}

func ex1() {
	input := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	result := search(input, target)
	fmt.Printf("Does the result [%v] match the expectation? [%v]\n", result, result == 4)
}

func ex2() {
	input := []int{3, 1}
	target := 3
	result := search(input, target)
	fmt.Printf("Does the result [%v] match the expectation? [%v]\n", result, result == 0)
}

// O(log n)
func search(nums []int, target int) int {

	// Start a binary search on it...
	var left int
	right := len(nums) - 1

	for right >= left {

		mid := (left + right) / 2

		// Return if the target is the mid-point
		if target == nums[mid] {
			return mid
		}

		if nums[left] <= nums[mid] { // left side is sorted

			if target >= nums[left] && target < nums[mid] {
				// target is in the left slice, move the right pointer to the top of this left slice
				right = mid - 1
			} else {
				// move the left pointer to the bottom of the right slice instead
				left = mid + 1
			}

		} else { // right side is sorted

			if target > nums[mid] && target <= nums[right] {
				// target is in right slice, move left pointer to bottom of this right slice
				left = mid + 1
			} else {
				// move the right pointer to the top of the left slice instead
				right = mid - 1
			}
		}
	}

	return -1
}

// O(n)
func searchBruteForce(nums []int, target int) int {

	// start with a brute force approach

	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		}
	}

	return -1
}

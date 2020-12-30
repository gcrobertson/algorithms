package main

import (
	"fmt"
	"math"
)

func main() {

	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	result := maxAreaBruteForce(input)

	fmt.Printf("Does the result %v match the expectation? %v\n", result, result == 49)

}

// Brute Force approach solves in O(n*n) time complexity.
func maxAreaBruteForce(height []int) int {

	var maxarea int

	// O(n)
	for i := 0; i < len(height); i++ {

		// O(n*n)
		for j := i + 1; j < len(height); j++ {

			minheight := int(math.Min(float64(height[i]), float64(height[j])))
			area := minheight * (j - i)
			maxarea = int(math.Max(float64(maxarea), float64(area)))
		}
	}

	return maxarea
}

// Solved in O(n) using 2 pointers
func maxArea(height []int) int {

	var maxarea int
	var left int
	right := len(height) - 1

	for left < right {

		h := int(math.Min(float64(height[left]), float64(height[right])))
		area := h * (right - left)
		maxarea = int(math.Max(float64(maxarea), float64(area)))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxarea
}

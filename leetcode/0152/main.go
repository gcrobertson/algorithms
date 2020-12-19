package main

import (
	"fmt"
	"math"
)

func main() {

	input1 := []int{2, 3, -2, 4}
	expect1 := 6

	input2 := []int{-2, 0, -1}
	expect2 := 0

	input3 := []int{-2, -5, 10, -3}
	expect3 := 150

	input4 := []int{0, 2}
	expect4 := 2

	input5 := []int{-2, 3, -4}
	expect5 := 24

	result1 := maxProduct(input1)
	result2 := maxProduct(input2)
	result3 := maxProduct(input3)
	result4 := maxProduct(input4)
	result5 := maxProduct(input5)

	fmt.Printf("Is result1 correct? %v\n", result1 == expect1)
	fmt.Printf("Is result2 correct? %v\n", result2 == expect2)
	fmt.Printf("Is result3 correct? %v\n", result3 == expect3)
	fmt.Printf("Is result4 correct? %v\n", result4 == expect4)
	fmt.Printf("Is result5 correct? %v\n", result5 == expect5)
}

// Approach 1: Brute Force
func maxProductBruteForce(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	max := nums[0]

	for i := 0; i < len(nums); i++ {

		product := 1

		for j := i; j < len(nums); j++ {

			product *= nums[j]

			if product > max {
				max = product
			}
		}
	}

	return max
}

// Approach 3: User Comment
func maxProductUserComment(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	max := nums[0]

	product := 1

	for _, num := range nums {
		product *= num
		max = int(math.Max(float64(max), float64(product)))

		if product == 0 {
			product = 1
		}
	}

	product = 1

	for i := len(nums) - 1; i > -1; i-- {
		product *= nums[i]
		max = int(math.Max(float64(max), float64(product)))
		if product == 0 {
			product = 1
		}
	}

	return max
}

// Approach 2: Dynamic Programming
func maxProductDynamicProgramming(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	max, min, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {

		curr := nums[i]

		tempMax := int(math.Max(float64(curr), math.Max(float64(max*curr), float64(min*curr))))

		min = int(math.Min(float64(curr), math.Min(float64(max*curr), float64(min*curr))))

		max = tempMax

		result = int(math.Max(float64(max), float64(result)))
	}

	return result
}

func maxProduct(nums []int) int {
	return maxProductDynamicProgramming(nums)
}

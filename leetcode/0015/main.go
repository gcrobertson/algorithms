package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {

	input := []int{-1, 0, 1, 2, -1, -4}

	// expect := [][]int{
	// 	[]int{-1, -1, 2},
	// 	[]int{-1, 0, 1},
	// }

	result := threeSum(input)

	fmt.Printf("The result: %v\n", result)
}

// Brute Force ... Time limit exceeded ...
// This is a working solution though.
func threeSumBruteForce(nums []int) [][]int {

	var result [][]int

	// Is there an easier brute force approach?
	// Yes, this seems to work but allows duplicates.
	// Sorting the temporary result in ascending order and checking the result set will prevent duplicates.

	for i := 0; i < len(nums)-2; i++ {

		num1 := nums[i]
		for j := i + 1; j < len(nums)-1; j++ {

			num2 := nums[j]
			for k := j + 1; k < len(nums); k++ {

				num3 := nums[k]
				if num1+num2+num3 == 0 {

					tempSlice := []int{num1, num2, num3}
					sort.Ints(tempSlice)

					var new = true
					for _, ix := range result {
						if reflect.DeepEqual(tempSlice, ix) {
							new = false
						}
					}
					if new == true {
						result = append(result, tempSlice)
					}
				}
			}
		}
	}

	return result
}

// Let me take a step back and go back over the following problems so that I can build off of their patterns:
// - https://leetcode.com/problems/two-sum
// - https://leetcode.com/problems/two-sum-ii-input-array-is-sorted

// Approach 1: 2 Pointers

// Approach 2: Hash set

// Approach 3: No sort

/*	Approach 1: 2 Pointers
 *
 *
 *	Accepted: 	32 ms	[faster than 71.97%]
 *				6.9 MB	[less than 85.36%]
 */
func threeSum(nums []int) [][]int {

	// O(n) Sort all of the inputs
	sort.Ints(nums)

	var result [][]int

	// O(n) Iterate over all values
	for i := 0; i < len(nums); i++ {

		if nums[i] > 0 {
			break // I think this is maintaining that 1 number is negative or all numbers must be 0
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue // Skip on any duplicates. This works because slice is ordered ..
		}

		// call TwoSum for the current position i
		result = twoSum(i, nums, result)
	}

	return result
}

/*  Approach 1: `2 Pointers`
 *
 */
func twoSum(i int, nums []int, result [][]int) [][]int {

	lo := i + 1
	hi := len(nums) - 1

	for lo < hi {
		sum := nums[lo] + nums[hi] + nums[i]
		if sum < 0 {
			lo++
		} else if sum == 0 {
			xi := []int{nums[lo], nums[i], nums[hi]}
			sort.Ints(xi)
			result = append(result, xi)
			hi--
			lo++
			for lo < hi && nums[lo] == nums[lo-1] {
				lo++
			}
		} else {
			hi--
		}
	}

	return result
}

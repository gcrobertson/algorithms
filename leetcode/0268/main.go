package main

import (
	"fmt"
	"sort"
)

func main() {

	ex1()
	ex2()
	ex3()
	ex4()
}

func ex1() {
	input := []int{3, 0, 1}
	expect := 2
	result := missingNumber(input)
	fmt.Printf("Does the result [%v] match our expectation? [%v]\n", result, result == expect)
}

func ex2() {
	input := []int{0, 1}
	expect := 2
	result := missingNumber(input)
	fmt.Printf("Does the result [%v] match our expectation? [%v]\n", result, result == expect)
}

func ex3() {
	input := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	expect := 8
	result := missingNumber(input)
	fmt.Printf("Does the result [%v] match our expectation? [%v]\n", result, result == expect)
}

func ex4() {
	input := []int{0}
	expect := 1
	result := missingNumber(input)
	fmt.Printf("Does the result [%v] match our expectation? [%v]\n", result, result == expect)
}

//  Approach 2: HashSet
//	Runtime: O(n + n + n)
//	Runtime: 20 ms, faster than 24.26% of Go online submissions for Missing Number.
//	Memory Usage: 6 MB, less than 98.53% of Go online submissions for Missing Number.
func missingNumberHashSet(nums []int) int {

	r := len(nums)
	max := r + 1
	hash := make(map[int]bool, max)

	for i := 0; i < max; i++ {
		hash[i] = false
	}
	for i := 0; i < len(nums); i++ {
		hash[nums[i]] = true
	}
	for i := 0; i < max; i++ {
		if hash[i] == false {
			return i
		}
	}
	return -1
}

// Approach 1: Sort
// Runtime: O(n + n)
// Runtime: 28 ms, faster than 10.66% of Go online submissions for Missing Number.
// Memory Usage: 6 MB, less than 98.53% of Go online submissions for Missing Number.
func missingNumberSort(nums []int) int {

	// O(n)
	sort.Ints(nums)

	// O(1)
	if nums[0] != 0 {
		return 0
	}

	// O(1)
	if nums[len(nums)-1] != len(nums) {
		return len(nums)
	}

	// O(n)
	for i := 1; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}

	return -1
}

// Approach 3: Bit Manipulation
// Runtime: O(n)
// Space: O(1)

// XOR is its own inverse -- what does this mean?
// Remember: If exclusively 1 of the input bits are 1, the resulting output is 1, otherwise the output is 0.
// 1 ^ 0 = 1
// 1 ^ 1 = 0
// 0 ^ 0 = 0
// 0 ^ 1 = 1

/*	missing = 4 ^ (0^0) ^ (1^1) ^ (2^3) ^ (3^4)
 *	missing = (4^4) ^ (0^0) ^ (1^1) ^ 2 ^ (3^3)
 *  missing = 0 ^ 0 ^ 0 ^ 2 ^ 0
 *  missing = 2
 */

// Runtime: 16 ms, faster than 86.40% of Go online submissions for Missing Number.
// Memory Usage: 6 MB, less than 98.53% of Go online submissions for Missing Number.
func missingNumberXOR(nums []int) int {

	missing := len(nums)

	for i := 0; i < len(nums); i++ {

		missing = missing ^ (i ^ nums[i])

	}

	return missing
}

// Approach 4: Gauss' Formula
// Runtime: 12 ms, faster than 98.90% of Go online submissions for Missing Number.
// Memory Usage: 6 MB, less than 98.53% of Go online submissions for Missing Number.
func missingNumber(nums []int) int {

	// Gauss' Formula can be calculated in O(1) time
	n := len(nums)
	gauss := (n * (n + 1)) / 2

	var sum int

	// O(n) loop through once to get total sum
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	// Return the difference as it is the missing number
	return gauss - sum
}

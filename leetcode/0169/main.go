package main

import (
	"fmt"
	"sort"
)

func main() {
	ex1()
	ex2()
}

func ex1() {
	input := []int{3, 2, 3}
	expect := 3
	result := majorityElement(input)
	fmt.Printf("result [%v] matches expectation [%v]?\n", result, result == expect)
}

func ex2() {
	input := []int{2, 2, 1, 1, 1, 2, 2}
	expect := 2
	result := majorityElement(input)
	fmt.Printf("result [%v] matches expectation [%v]?\n", result, result == expect)
}

// Runtime: 20 ms, faster than 61.09% of Go online submissions for Majority Element.
// Memory Usage: 6 MB, less than 61.40% of Go online submissions for Majority Element.
func majorityElementUsingHash(nums []int) int {
	majority := len(nums)/2 + 1
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hash[nums[i]]++
		if hash[nums[i]] == majority {
			return nums[i]
		}
	}
	return -1
}

// Runtime: 20 ms, faster than 61.09% of Go online submissions for Majority Element.
// Memory Usage: 5.9 MB, less than 61.40% of Go online submissions for Majority Element.
// Sorting.  Because it is `majority` and not `most common element` this works ...
func majorityElementSort(nums []int) int {

	// O(log n)
	sort.Ints(nums)

	// O(1)
	return nums[(len(nums)-1)/2]
}

// Boyer-Moore Voting Algorithm
// O(n) time
// O(1) space
// Runtime: 16 ms, faster than 95.44% of Go online submissions for Majority Element.
// Memory Usage: 5.9 MB, less than 61.40% of Go online submissions for Majority Element.
func majorityElement(nums []int) int {

	var count, candidate int

	for _, num := range nums {

		if count == 0 {
			candidate = num
		}

		if candidate == num {
			count++
		} else {
			count--
		}
	}

	return candidate
}

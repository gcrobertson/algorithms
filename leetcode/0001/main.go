package main

import (
	"fmt"
	"sort"
)

func main() {
	ex1()
	// ex2()
	// ex3()
	// ex4()
}

func ex1() {
	nums := []int{2, 7, 11, 15}
	sum := 9
	result := twoSum(nums, sum)
	fmt.Printf("ex1(): expect: %v\t result: %v\n", []int{0, 1}, result)
}

func ex2() {
	nums := []int{2, 3, 4}
	target := 6
	result := twoSum(nums, target)
	fmt.Printf("ex2(): expect: %v\t result: %v\n", []int{0, 2}, result)
}

func ex3() {
	nums := []int{3, 2, 4}
	target := 6
	result := twoSum(nums, target)
	fmt.Printf("ex3(): expect: %v\t result: %v\n", []int{1, 2}, result)
}

func ex4() {
	nums := []int{2, 5, 5, 11}
	target := 10
	result := twoSum(nums, target)
	fmt.Printf("ex4(): expect: %v\t result: %v\n", []int{1, 2}, result)
}

/*
 *	Given an array of integers, return indices of the two numbers
 *	such that they add up to a specific target.
 *
 *	You may assume that each input would have exactly 1 solution, and you may not use the same element twice.
 *
 *	Example:
 *
 *	Given nums: [2,7,11,15], target: 9
 *
 *	Because nums[0] & nums[1] = 2 + 7 = 9,
 *
 *	return [0,1]
 *
 *
 */

// https://leetcode.com/list/xy8qu5y1/ google top 50: #1
// https://leetcode.com/problems/two-sum/

// Brute Force approach ... O(n²)
func twoSumBruteForce(nums []int, target int) []int {

	for i := 0; i < len(nums)-1; i++ {

		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				return []int{i + 1, j + 1} // non-zero based index result, so add 1 to each index ...
			}
		}
	}
	return []int{}
}

func twoSumTwoPass(nums []int, target int) []int {

	hash := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		hash[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		search := target - nums[i]
		if _, ok := hash[search]; ok && hash[search] != i {
			return []int{i, hash[search]}
		}
	}

	return []int{}
}

// Also O(n) but not O(n+n) as previous ...
func twoSum(nums []int, target int) []int {

	hash := make(map[int]int) // `value` : `index`

	for i := 0; i < len(nums); i++ {

		search := target - nums[i]

		//fmt.Printf("hash: %v\n", hash)
		//fmt.Printf("index: %v\t value: %v\t search: %v\n", i, nums[i], search)

		if _, ok := hash[search]; ok {
			x := []int{i, hash[search]}
			sort.Ints(x)
			return x
		}

		hash[nums[i]] = i // `value` : `index`
	}

	return []int{}
}

package main

import "fmt"

func main() {

	nums := []int{2, 7, 11, 15}
	sum := 9

	fmt.Printf("%+v", twoSum(nums, sum))
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

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums)-1; i++ {

		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

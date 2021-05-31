package main

import (
	"fmt"
	"reflect"
)

func main() {
	ex1()
	ex2()
	ex3()
}

func ex1() {
	input := []int{8, 1, 2, 2, 3}
	expect := []int{4, 0, 1, 1, 3}
	result := smallerNumbersThanCurrent(input)
	fmt.Printf("Example 1? %v. Expect: %v Result: %v\n", reflect.DeepEqual(expect, result), expect, result)
}

func ex2() {
	input := []int{6, 5, 4, 8}
	expect := []int{2, 1, 0, 3}
	result := smallerNumbersThanCurrent(input)
	fmt.Printf("Example 2? %v. Expect: %v Result: %v\n", reflect.DeepEqual(expect, result), expect, result)
}

func ex3() {
	input := []int{7, 7, 7, 7}
	expect := []int{0, 0, 0, 0}
	result := smallerNumbersThanCurrent(input)
	fmt.Printf("Example 3? %v. Expect: %v Result: %v\n", reflect.DeepEqual(expect, result), expect, result)
}

// Input constraints:
// 2 <= nums.length <= 500
// 0 <= nums[i] <= 100
func smallerNumbersThanCurrent(nums []int) []int {

	var bucket [101]int

	for i := 0; i < len(nums); i++ {
		bucket[nums[i]]++
	}

	for i := 1; i < len(bucket); i++ {
		bucket[i] += bucket[i-1]
	}

	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		result[i] = bucket[nums[i]-1]
	}

	return result
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	ex1()
	// ex2()
	// ex3()
	// ex4()
}

func ex1() {
	input := []int{0, 1, 0, 3, 12}
	expect := []int{1, 3, 12, 0, 0}
	moveZeroes(input)
	fmt.Printf("result [%v] matches expectation[%v] ? [%v]\n", input, expect, reflect.DeepEqual(input, expect))
}

func ex2() {
	input := []int{0}
	expect := []int{0}
	moveZeroes(input)
	fmt.Printf("result [%v] matches expectation[%v] ? [%v]\n", input, expect, reflect.DeepEqual(input, expect))
}

func ex3() {
	input := []int{0, 0}
	expect := []int{0, 0}
	moveZeroes(input)
	fmt.Printf("result [%v] matches expectation[%v] ? [%v]\n", input, expect, reflect.DeepEqual(input, expect))
}

func ex4() {
	input := []int{0, 0, 1}
	expect := []int{1, 0, 0}
	moveZeroes(input)
	fmt.Printf("result [%v] matches expectation[%v] ? [%v]\n", input, expect, reflect.DeepEqual(input, expect))
}

// Runtime: 4 ms, faster than 91.68% of Go online submissions for Move Zeroes.
// Memory Usage: 3.8 MB, less than 19.41% of Go online submissions for Move Zeroes.
func moveZeroesSolution1(nums []int) {

	if len(nums) < 2 {
		return
	}

	var p1 int
	var p2 = len(nums) - 1

	for nums[p2] == 0 && p2 > 0 {
		p2--
	}

	if p2 == p1 {
		return
	}

	// O(n)
	for p1 < len(nums) && p1 != p2 {

		if nums[p1] != 0 {
			p1++
			continue
		}

		nums = append(nums[0:p1], append(nums[p1+1:], 0)...)
		p2--
	}
}

// Time complexity O(n)
// Space complexity O(1)
// Total # of writes is `n`
// Runtime: 4 ms, faster than 91.68% of Go online submissions for Move Zeroes.
// Memory Usage: 3.8 MB, less than 19.41% of Go online submissions for Move Zeroes.
func moveZeroesSolution2(nums []int) {

	var idx int

	// O(n) to preserve order of all non-zero values. will be 1,3,12,3,12 with idx = 3
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[idx] = nums[i]
			idx++
		}
	}

	fmt.Println(nums, idx)

	// O(n)
	for i := idx; i < len(nums); i++ {
		nums[i] = 0
	}

}

// Runtime: 4 ms, faster than 91.68% of Go online submissions for Move Zeroes.
// Memory Usage: 3.8 MB, less than 19.41% of Go online submissions for Move Zeroes.
func moveZeroes(nums []int) {
	var p1, p2 int
	// O(n)
	for p2 < len(nums) {
		if nums[p2] != 0 {
			nums[p1], nums[p2] = nums[p2], nums[p1]
			p1++
		}
		p2++
	}
}

package main

import (
	"fmt"
	"reflect"
)

/*
 *
 *
 *
 */
func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
}

func ex1() {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expect := []int{1, 4, 2, 7, 5, 3, 8, 6, 9}
	result := findDiagonalOrder(input)
	fmt.Printf("result [%v] matches expectation? [%v]\n", result, reflect.DeepEqual(result, expect))
}

func ex2() {
	input := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7},
		{8},
		{9, 10, 11},
		{12, 13, 14, 15, 16},
	}
	expect := []int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16}
	result := findDiagonalOrder(input)
	fmt.Printf("result [%v] matches expectation? [%v]\n", result, reflect.DeepEqual(result, expect))
}

func ex3() {
	input := [][]int{
		{1, 2, 3},
		{4},
		{5, 6, 7},
		{8},
		{9, 10, 11},
	}
	expect := []int{1, 4, 2, 5, 3, 8, 6, 9, 7, 10, 11}
	result := findDiagonalOrder(input)
	fmt.Printf("result [%v] matches expectation? [%v]\n", result, reflect.DeepEqual(result, expect))
}

func ex4() {
	input := [][]int{
		{1, 2, 3, 4, 5, 6},
	}
	expect := []int{1, 2, 3, 4, 5, 6}
	result := findDiagonalOrder(input)
	fmt.Printf("result [%v] matches expectation? [%v]\n", result, reflect.DeepEqual(result, expect))
}

func ex5() {
	input := [][]int{
		{14, 12, 19, 16, 9},
		{13, 14, 15, 8, 11},
		{11, 13, 1},
	}
	expect := []int{14, 13, 12, 11, 14, 19, 13, 15, 16, 1, 8, 9, 11}
	result := findDiagonalOrder(input)
	fmt.Printf("result [%v] matches expectation? [%v]\n", result, reflect.DeepEqual(result, expect))
}

// This works but time limited exceeded. Need an alternative pattern ...
func findDiagonalOrderTimeLimitExceeded(nums [][]int) []int {
	var result []int
	if len(nums) == 0 {
		return result
	}

	var maxrow int // necessary to store max row length so that ex5() can be solved ...

	for y := 0; y < len(nums); y++ {
		result = append(result, nums[y][0])
		// fmt.Printf("starting coordinate: [y,x] = [%v,0] = %v\n", y, nums[y][0])
		y1 := y - 1
		x1 := 1

		for y1 > -1 {
			if x1 > len(nums[y1])-1 {
				x1++
				y1--
				continue
			}
			// fmt.Printf("new coordinate: [y,x] = [%v,%v] = %v\n", y1, x1, nums[y1][x1])
			result = append(result, nums[y1][x1])
			x1++
			y1--
		}

		if len(nums[y]) > maxrow {
			maxrow = len(nums[y])
		}

		if y == len(nums)-1 {
			y2 := y
			x2 := 1
			// loop needs to be based horizontally this time with the vertical one in it...! confusing.
			for x2 < maxrow {
				y2 = y
				x3 := x2
				for y2 > -1 {
					if x3 > len(nums[y2])-1 {
						x3++
						y2--
						continue
					}
					// fmt.Printf("Last one ... coordinate: [y,x] = [%v,%v]\n", y2, x2)
					result = append(result, nums[y2][x3])
					x3++
					y2--
				}
				x2++
			}
		}

	}

	return result
}

// Runtime: 180 ms, faster than 94.12% of Go online submissions for Diagonal Traverse II.
// Memory Usage: 22.3 MB, less than 20.59% of Go online submissions for Diagonal Traverse II.
func findDiagonalOrder(nums [][]int) []int {

	var result []int
	if len(nums) == 0 {
		return result

	}
	var maxKey int // maximum key inserted into the map, i.e.: max value of i+j indices

	buckets := make(map[int][]int, len(nums))

	// go through all cells in matrix, put values buckets. for example, bucket([y,x]) = {[2,0], [1,1], [0,2]}
	// O(n*m)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			buckets[i+j] = append(buckets[i+j], nums[i][j])
			if i+j > maxKey {
				maxKey = i + j
			}
		}
	}

	//O((n+m-2) * ?) ? = max len of a bucket
	for i := 0; i <= maxKey; i++ {

		bucket := buckets[i]

		for x := len(bucket) - 1; x > -1; x-- {
			result = append(result, bucket[x])
		}
	}

	return result
}

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
		{13, 114, 15, 8, 11}, // @TODO: It is not finding the 11 because the final loop uses the length of the last row
		{11, 13, 1},
	}
	expect := []int{14, 13, 12, 11, 14, 19, 13, 15, 16, 1, 8, 9, 11}
	result := findDiagonalOrder(input)
	fmt.Printf("result [%v] matches expectation? [%v]\n", result, reflect.DeepEqual(result, expect))
}

/*
 *
 *
 *
 */
func findDiagonalOrder(nums [][]int) []int {
	var result []int
	if len(nums) == 0 {
		return result
	}

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

		if y == len(nums)-1 {
			y2 := y
			x2 := 1
			// loop needs to be based horizontally this time with the vertical one in it...! confusing.
			for x2 < len(nums[y]) {
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

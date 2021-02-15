package main

import (
	"fmt"
	"sort"
)

func main() {

	// i := []int{1, 2, 2, 3, 3, 3}
	// fmt.Printf("%v\n", findLucky(i) == 3)

	i := []int{2, 2, 3, 4}
	fmt.Printf("%v == %v? %v\n", findLucky(i), 2, findLucky(i) == 2)
}

// Runtime: 4 ms, faster than 91.67% of Go online submissions for Find Lucky Integer in an Array.
// Memory Usage: 3 MB, less than 94.44% of Go online submissions for Find Lucky Integer in an Array.
func findLucky(arr []int) int {

	// O(log n)
	sort.Ints(arr)

	// O(1)
	var c int

	// O(n)
	for i := len(arr) - 1; i >= 0; i-- {

		// O(1)
		c++

		// O(1)
		if i == 0 || arr[i] != arr[i-1] {

			// O(1)
			if arr[i] == c {
				return c
			}

			// O(1)
			c = 0
		}

	}

	// O(1)
	return -1
}

// Runtime: 4 ms, faster than 91.67% of Go online submissions for Find Lucky Integer in an Array.
// Memory Usage: 3.1 MB, less than 33.33% of Go online submissions for Find Lucky Integer in an Array.
/*
 *
 *
 *
 */
func findLuckyHash(arr []int) int {

	if len(arr) == 0 {
		return -1
	}

	m := make(map[int]int)

	for _, v := range arr {
		m[v]++
	}

	result := -1

	for i, v := range m {
		if i == v && v > result {
			result = v
		}
	}

	return result
}

// Runtime: 4 ms, faster than 91.67% of Go online submissions for Find Lucky Integer in an Array.
// Memory Usage: 3.1 MB, less than 33.33% of Go online submissions for Find Lucky Integer in an Array.
/*
 *
 *
 *
 */
func findLuckyArray(arr []int) int {

	if len(arr) == 0 {
		return -1
	}

	m := make([]int, 501)
	for _, v := range arr {
		m[v]++
	}

	for i := 500; i > 0; i-- {
		if i == m[i] {
			return i
		}
	}

	return -1
}

package main

import (
	"fmt"
	"reflect"
)

func main() {

	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expect := []int{1, 2, 4, 7, 5, 3, 6, 8, 9}
	result := findDiagonalOrder(input)

	fmt.Printf("result [%v] matches expectation? [%v]\n", result, reflect.DeepEqual(expect, result))
}

// Runtime: 32 ms, faster than 28.90% of Go online submissions for Diagonal Traverse.
// Memory Usage: 7.8 MB, less than 5.65% of Go online submissions for Diagonal Traverse.
func findDiagonalOrder(matrix [][]int) []int {

	if len(matrix) == 0 {
		return []int{}
	}

	// step 1 put into buckets
	buckets := make(map[int][]int)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {

			buckets[i+j] = append(buckets[i+j], matrix[i][j])
		}
	}

	var result []int
	for i := 0; i < len(buckets); i++ {

		if i%2 == 0 {
			// reverse order for an even matrix[i] bucket.
			for j, k := 0, len(buckets[i])-1; j < k; j, k = j+1, k-1 {
				buckets[i][j], buckets[i][k] = buckets[i][k], buckets[i][j]
			}
		}
		result = append(result, buckets[i]...)
	}

	return result
}

// Runtime: 32 ms, faster than 28.90% of Go online submissions for Diagonal Traverse.
// Memory Usage: 7.8 MB, less than 5.65% of Go online submissions for Diagonal Traverse.
func findDiagonalOrder2(matrix [][]int) []int {

	if len(matrix) == 0 {
		return []int{}
	}

	// step 1 put into buckets
	buckets := make(map[int][]int)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {

			buckets[i+j] = append(buckets[i+j], matrix[i][j])
		}
	}

	// now just figure out how to output them.
	// how are they all stored now? all are stored top to bottom...
	// fmt.Println(buckets)

	var result []int
	for i := 0; i < len(buckets); i++ {

		if i%2 == 0 {
			// on even matrix[i] coordinates we go bottom to top,
			for j := len(buckets[i]) - 1; j > -1; j-- {
				result = append(result, buckets[i][j])
			}
			// on odd  matrix[i] coordinates we go top to bottom,
		} else {
			result = append(result, buckets[i]...)
		}

	}

	return result
}

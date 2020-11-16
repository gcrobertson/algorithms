package main

import (
	"fmt"
	"math/rand"
	"time"
)

// w = [1, 3]
// 0 = 1,
// 1 = 3,
// 0, 1, 1, 1
func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	w := []int{3, 14, 1, 7}
	obj := Constructor(w)

	// fmt.Printf("%v", obj)

	result := obj.PickIndexLinearSearch()
	result2 := obj.PickIndexBinarySearch()

	fmt.Printf("Picked index: %v\n", result)
	fmt.Printf("Picked index: %v\n", result2)

}

// Solution ...
type Solution struct {
	prefixSums []int
	totalSum   int
}

// Constructor ...
func Constructor(w []int) Solution {

	prefixSums := make([]int, len(w))
	var prefixSum int

	for i := 0; i < len(w); i++ {
		prefixSum += w[i]
		prefixSums[i] = prefixSum
	}

	return Solution{
		prefixSums: prefixSums,
		totalSum:   prefixSum,
	}

}

// PickIndexBinarySearch ...
func (sol *Solution) PickIndexBinarySearch() int {

	target := rand.Intn(sol.totalSum)
	var low int
	high := len(sol.prefixSums)

	for low < high {
		mid := (high - low) / 2

		if target < sol.prefixSums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

// PickIndexLinearSearch ... using a linear search
func (sol *Solution) PickIndexLinearSearch() int {

	target := rand.Intn(sol.totalSum)
	var index int

	for i := 0; i < len(sol.prefixSums); i++ {
		if target < sol.prefixSums[i] {
			index = i
			break
		}
	}

	return index
}

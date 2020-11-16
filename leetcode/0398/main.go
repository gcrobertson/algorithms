package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	input := []int{1, 2, 3, 3, 3}
	obj := Constructor(input)
	value := 3
	param1 := obj.Pick(value)

	fmt.Printf("The index randomly selected for value [%v] is [%v]\n", value, param1)
}

// Solution ...
type Solution struct {
	nums []int
}

// Constructor ...
func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

// Pick ...
func (sol *Solution) Pick(target int) int {

	var count, index int
	for i := 0; i < len(sol.nums); i++ {
		if sol.nums[i] == target {
			count++
			if rand.Intn(count) == 0 { // we pick current number with probability 1/count [resevoir sampling]
				index = i
			}
		}
	}

	return index
}

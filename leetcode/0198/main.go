package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{1, 2, 3, 1}
	input2 := []int{2, 7, 9, 3, 1}

	fmt.Printf("max from input 1 should be 4: %v\n", rob(input))
	fmt.Printf("max from input 1 should be 12: %v\n", rob(input2))
	fmt.Printf("max from input 1 should be 4: %v\n", rob2(input))
	fmt.Printf("max from input 1 should be 12: %v\n", rob2(input2))
}

func rob2(nums []int) int {
	var prevMax, currMax int
	for _, x := range nums {
		var temp = currMax
		currMax = int(math.Max(float64(prevMax+x), float64(currMax)))
		prevMax = temp
	}
	return currMax
}

func rob(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	burgleMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		max := nums[i]
		for j := i - 2; j >= 0; j-- {
			if nums[i]+burgleMap[j] > max {
				max = nums[i] + burgleMap[j]
			}
		}
		burgleMap[i] = max
	}
	var result int
	for _, v := range burgleMap {
		if v > result {
			result = v
		}
	}
	return result
}

package main

import (
	"fmt"
	"math"
)

func rob(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}

	var1 := robber(nums[0 : len(nums)-1])
	var2 := robber(nums[1:len(nums)])

	return int(math.Max(float64(var1), float64(var2)))
}

func robber(nums []int) int {

	fmt.Printf("%v\n", nums)

	var prevMax, currMax int

	for _, x := range nums {

		var temp = currMax
		currMax = int(math.Max(float64(x+prevMax), float64(currMax)))
		prevMax = temp
	}

	return currMax
}

func main() {

}

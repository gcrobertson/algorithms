package main

import "fmt"

func main() {

	given := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	result := 5

	fmt.Printf("%v == %v?\n", result, removeDuplicates(given))
}

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	var p1 int
	for p2 := 1; p2 < len(nums); p2++ {
		if nums[p1] != nums[p2] {
			p1++
			nums[p1] = nums[p2]
		}
	}
	return p1 + 1

	/*
		i := 0

		for i < len(nums)-1 {
			if nums[i] == nums[i+1] {
				copy(nums[i:], nums[i+1:])
				nums[len(nums)-1] = 0
				nums = nums[:len(nums)-1]
			} else {
				i++
			}
		}

		return len(nums)
	*/
}

// func removeDuplicates(nums []int) int {
// 	i := 0
// 	r := len(nums)

// 	for i < r-1 {
// 		if nums[i] == nums[i+1] {
// 			nums = append(nums[0:i], nums[i+1:]...)
// 			r = len(nums)
// 		} else {
// 			i++
// 		}
// 	}

// 	return len(nums)
// }

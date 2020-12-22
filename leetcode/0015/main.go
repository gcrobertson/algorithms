package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {

	input := []int{-1, 0, 1, 2, -1, -4}

	// expect := [][]int{
	// 	[]int{-1, -1, 2},
	// 	[]int{-1, 0, 1},
	// }

	result := threeSum(input)

	fmt.Printf("The result: %v\n", result)
}

// Brute Force ... Time limit exceeded ...
// This is a working solution though.
func threeSumBruteForce(nums []int) [][]int {

	var result [][]int

	// Is there an easier brute force approach?
	// Yes, this seems to work but allows duplicates.
	// Sorting the temporary result in ascending order and checking the result set will prevent duplicates.

	for i := 0; i < len(nums)-2; i++ {

		num1 := nums[i]
		for j := i + 1; j < len(nums)-1; j++ {

			num2 := nums[j]
			for k := j + 1; k < len(nums); k++ {

				num3 := nums[k]
				if num1+num2+num3 == 0 {

					tempSlice := []int{num1, num2, num3}
					sort.Ints(tempSlice)

					var new = true
					for _, ix := range result {
						if reflect.DeepEqual(tempSlice, ix) {
							new = false
						}
					}
					if new == true {
						result = append(result, tempSlice)
					}
				}
			}
		}
	}

	return result
}

/*	What would make the Brute Force solution quicker?
 *
 *	The solution is re-testing a lot of the combinations.  How could I prevent that from happening?
 *
 *
 */
func threeSum(nums []int) [][]int {

	var result [][]int

	if len(nums) < 3 {
		return result
	}

	// Let me take a step back and go back over the following problems so that I can build off of their patterns:
	// - https://leetcode.com/problems/two-sum
	// - https://leetcode.com/problems/two-sum-ii-input-array-is-sorted

	return result
}

/*
	// We are given an array of ints,
	// If any 3 values in the array sum to 0, add them to our result if they are not already there ...

	// My first impression is I could collect all of the unique numbers in the array in a map along with the count of that number:
	// map[-4] = 1
	// map[-1] = 2
	// map[0]  = 1
	// map[1]  = 1
	// map[2]  = 1
	//
	// this could be done in O(n) time.

	// Then, I could form doubles or pairs with each unique combination.
	// map[-4] > [-1] = needs a +5
	// map[-4] > [ 0] = needs a +4
	// map[-4] > [ 1] = needs a +3
	// map[-4] > [ 2] = needs a +2, there is only 1.

	// At this point I could discard the -4.

	// map[-1] > [-1] = needs a +2, it is there!
	// map[-1] > [ 0] = needs a +1, it is there!
	// map[-1] > [ 1] = needs a  0, it is there!		// If I have any 2 already in the result set, I can just continue... because the 3rd number will always be the same.

	// This is not a quick solution but I think it will work...

		hashCount := make(map[int]int)

		// O(n)
		for _, num := range nums {
			hashCount[num]++
		}

*/

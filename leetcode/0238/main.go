package main

func main() {
}

func productExceptSelf(nums []int) []int {

	// Product of Array Except Self
	//
	// n integers,
	// n > 1
	// return output
	// output[i] =

	// input    = [1,2,3,4]
	// output   = [24,12,8,6]
	// [2*3*4, 1*3*4, 1*2*4, 1*2*3]

	// product := 1*2*3*4
	// solution for any key = product / key

	// O(1)
	product := 1
	zeroCount := 0

	// O(n)
	for _, v := range nums {
		if v == 0 {
			zeroCount++
		} else {
			product *= v
		}
	}

	result := make([]int, len(nums))

	if zeroCount > 1 {
		return result
	}

	for k, v := range nums {

		if v == 0 {
			result[k] = product
		} else if v != 0 && zeroCount == 1 {
			result[k] = 0
		} else {
			result[k] = product / v
		}

	}

	return result
}

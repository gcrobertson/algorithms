package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {

	ex1()
}

func ex1() {

	input := []int{1, 1, 2, 2, 2, 3}

	output := []int{3, 1, 1, 2, 2, 2}

	result := frequencySort(input)

	fmt.Printf("result [%v] is the expectation [%v]? [%v]\n", result, output, reflect.DeepEqual(output, result))
}

func frequencySort(nums []int) []int {
	tmp := make([]int, 201)
	for _, v := range nums {
		tmp[v+100]++
	}

	sort.Slice(nums, func(a, b int) bool {
		if tmp[nums[a]+100] == tmp[nums[b]+100] {
			return nums[a] > nums[b]
		}
		return tmp[nums[a]+100] < tmp[nums[b]+100]
	})
	return nums
}

/*
func frequencySort(nums []int) []int {

	h := make(map[int]int)

	max := nums[0]

	for _, v := range nums {
		h[v]++
		if v > max {
			max = v
		}
	}

	var result []int

	for val, num := range h {

		for i := 0; i < num; i++ {
			result = append(result, val)
		}
	}

	sort.Ints(result)

	return result
}
*/

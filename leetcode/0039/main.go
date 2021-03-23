package main

import (
	"fmt"
	"reflect"
	"sort"
)

func ex1() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	expect := [][]int{
		{2, 2, 3},
		{7},
	}
	result := combinationSum(candidates, target)
	fmt.Printf("Result [%v] is expected? [%v]\n", result, reflect.DeepEqual(expect, result))
}

func ex2() {
	candidates := []int{2, 3, 5}
	target := 8
	expect := [][]int{
		{2, 2, 2, 2},
		{2, 3, 3},
		{3, 5},
	}
	result := combinationSum(candidates, target)
	fmt.Printf("Result [%v] is expected? [%v]\n", result, reflect.DeepEqual(expect, result))
}

func ex3() {
	candidates := []int{2}
	target := 1
	var expect [][]int
	result := combinationSum(candidates, target)
	fmt.Printf("Result [%v] is expected? [%v]\n", result, reflect.DeepEqual(expect, result))
}

func ex4() {
	candidates := []int{2, 3, 5}
	target := 8
	expect := [][]int{
		{2, 2, 2, 2},
		{2, 3, 3},
		{3, 5},
	}
	result := combinationSum(candidates, target)
	fmt.Printf("Result [%v] is expected? [%v]\n", result, reflect.DeepEqual(expect, result))
}

func ex5() {
	candidates := []int{1}
	target := 1
	expect := [][]int{
		{1},
	}
	result := combinationSum(candidates, target)
	fmt.Printf("Result [%v] is expected? [%v]\n", result, reflect.DeepEqual(expect, result))
}

func main() {

	ex1()
	// ex2()
	// ex3()
	// ex4()
	// ex5()

}

// Attempt 1 fails for example 1. It works for 2,3,4,5
func combinationSumAttempt1(candidates []int, target int) [][]int {

	var result [][]int

	// create hash map
	hash := make(map[int]bool, len(candidates))
	for _, v := range candidates {
		hash[v] = true
	}

	for _, v := range candidates {

		// [2,2,2,2] = 8
		if target%v == 0 {
			t := []int{}
			for n := target / v; n > 0; n-- {
				t = append(t, v)
			}
			result = append(result, t)
		} else {

			n := target
			c := 0
			for {
				if n-v > 0 {
					n -= v
					c++
				} else {
					break
				}
			}

			if _, ok := hash[n]; ok {
				t := []int{}
				for c > 0 {
					t = append(t, v)
					c--
				}
				t = append(t, n)
				sort.Ints(t)
				result = append(result, t)
			}
		}
	}

	return result
}

// https://leetcode.com/problems/combination-sum/discuss/16541/Golang-solution-(6ms)
func combinationSum(candidates []int, target int) [][]int {

	var result [][]int
	sort.Ints(candidates)
	recursion(candidates, []int{}, target, 0, 0, &result)
	return result
}

func recursion(candidates []int, subset []int, target int, startIndex int, sum int, result *[][]int) {

	if sum == target {
		*result = append(*result, append([]int{}, subset...))
		return
	}
	for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {
		recursion(candidates, append(subset, candidates[i]), target, i, sum+candidates[i], result)
	}
}

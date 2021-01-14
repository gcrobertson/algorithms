package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// ex1()
	// ex2()
	// ex3()
	// ex4()
	// ex5()
	ex6()
	fmt.Printf("Total iterations: %v\n", totalIterations)
}

func ex1() {
	coins := []int{1, 2, 5}
	amount := 11
	expect := 3 // 5 + 5 + 1
	result := coinChange(coins, amount)
	fmt.Printf("Does result [%v] match expectation? [%v]\n", result, result == expect)
}
func ex2() {
	coins := []int{2}
	amount := 3
	expect := -1 // not possible
	result := coinChange(coins, amount)
	fmt.Printf("Does result [%v] match expectation? [%v]\n", result, result == expect)
}
func ex3() {
	coins := []int{1}
	amount := 0
	expect := 0 // no coins
	result := coinChange(coins, amount)
	fmt.Printf("Does result [%v] match expectation? [%v]\n", result, result == expect)
}
func ex4() {
	coins := []int{1}
	amount := 1
	expect := 1 // 1
	result := coinChange(coins, amount)
	fmt.Printf("Does result [%v] match expectation? [%v]\n", result, result == expect)
}
func ex5() {
	coins := []int{1}
	amount := 2
	expect := 2 // 1 + 1
	result := coinChange(coins, amount)
	fmt.Printf("Does result [%v] match expectation? [%v]\n", result, result == expect)
}

func ex6() {
	coins := []int{429, 171, 485, 26, 381, 31, 290}
	amount := 8440
	expect := 20
	result := coinChange(coins, amount)
	fmt.Printf("Does result [%v] match expectation? [%v]\n", result, result == expect)
}

var res int
var totalIterations int

func coinChange(coins []int, amount int) int {
	res = math.MaxInt32
	sort.Ints(coins)
	helper(coins, amount, 0, 0)
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

// coins	= original []int containing all coins
// remain	= amount remaining
// count	= current count of coins used
// coinIdx  = index of current coins slice
func helper(coins []int, remainingAmount int, coinCount int, coinIdx int) {
	if remainingAmount < 0 {
		return
	}
	if remainingAmount == 0 {
		if coinCount < res {
			res = coinCount
		}
		return
	}
	for i := coinIdx; i < len(coins); i++ {
		if coins[i] <= remainingAmount && remainingAmount < coins[i]*(res-coinCount) { // still do not understand: && remainingAmount < coins[i]*(res-coinCount)
			totalIterations++
			helper(coins, remainingAmount-coins[i], coinCount+1, i)
		}
	}
}

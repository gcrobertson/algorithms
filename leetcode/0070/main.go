package main

import "fmt"

// O(n)
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
// Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Climbing Stairs.
func climbStairs(n int) int {

	// NOTE: I knew as long as there were at least 2 steps to take, each step would have both +1 and +2
	// The key to solving this as to physically write out the first 5 steps, where I noticed the Fibonacci sequence
	// Each subsequent step was just the sum of the previous 2 steps ...

	if n == 1 {
		return n
	}

	var prevCount int
	currentStep, currentCount := 1, 1

	for currentStep <= n {
		t := currentCount
		currentCount = currentCount + prevCount
		prevCount = t
		currentStep++
	}

	return currentCount
}

func main() {

	// You are climbing a staircase. It takes n steps to reach the top.
	// Each time you can either climb 1 or 2 steps.
	// In how many distinct ways can you climb to the top?
	// constraint:
	// 1 <= n <= 45

	ex1()
	ex2()
	ex3()
}

func ex1() {

	input := 2
	// 1 + 1
	// 2
	expect := 2
	result := climbStairs(input)

	fmt.Printf("Does result [%v] match expectation? [%v]\n", result, result == expect)
}

func ex2() {

	input := 3
	// 1 + 1 + 1
	// 1 + 2
	// 2 + 1
	expect := 3
	result := climbStairs(input)

	fmt.Printf("Does result [%v] match expectation? [%v]\n", result, result == expect)
}

func ex3() {

	input := 4
	// 1 + 1 + 1
	// 1 + 2
	// 2 + 1
	expect := 6
	result := climbStairs(input)

	fmt.Printf("Does result [%v] match expectation? [%v]\n", result, result == expect)
}

/*


type treeNode struct {
	stairsLeft  int
	currentStep int
	nextOne     *treeNode
	nextTwo     *treeNode
}


*/

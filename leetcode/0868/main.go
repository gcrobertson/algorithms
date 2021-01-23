package main

import (
	"fmt"
	"math"
)

func main() {

	ex1()
	// ex2()
	// ex3()
	// ex4()
	// ex5()
	// ex6()
}

func ex1() {
	input := 22
	expect := 2
	result := binaryGap(input)
	fmt.Printf("Does result [%v] match expectation? [%v]\n", result, result == expect)
}

func ex2() {
	input := 5
	expect := 2
	result := binaryGap(input)
	fmt.Printf("Does result [%v] match expectation? [%v]\n", result, result == expect)
}

func ex3() {
	input := 6
	expect := 1
	result := binaryGap(input)
	fmt.Printf("Does result [%v] match expectation? [%v]\n", result, result == expect)
}

func ex4() {
	input := 8
	expect := 0
	result := binaryGap(input)
	fmt.Printf("Does result [%v] match expectation? [%v]\n", result, result == expect)
}

func ex5() {
	input := 1
	expect := 0
	result := binaryGap(input)
	fmt.Printf("Does result [%v] match expectation? [%v]\n", result, result == expect)
}

func ex6() {
	input := 13
	expect := 2
	result := binaryGap(input)
	fmt.Printf("Does result [%v] match expectation? [%v]\n", result, result == expect)
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Gap.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Binary Gap.
func binaryGapLinearTime(n int) int {

	bits := fmt.Sprintf("%b", n)

	k1 := -1
	k2 := -1
	max := -1

	for idx, bit := range bits {

		if bit == '1' {
			if k1 == -1 {
				k1 = idx
			} else if k2 == -1 {
				k2 = idx

				if k2-k1 > max {
					max = k2 - k1
				}

				k1 = k2
				k2 = -1
			}
		}

	}

	if max > 0 {
		return max
	}

	return 0
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Gap.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Binary Gap.
func binaryGap(N int) int {

	var max, current int

	/*
		before: 10110   after: 1011
		before: 1011    after: 101
		before: 101     after: 10
		before: 10      after: 1
		before: 1       after: 0
	*/

	for N > 0 {

		if N&1 == 1 { // Check if right most bit is a one

			max = int(math.Max(float64(max), float64(current)))
			current = 1

		} else if current > 0 {
			current++ // Increase any existing gap
		}

		N = N >> 1 // Right shift by 1

		fmt.Printf("after: %b\n", N)
	}

	return max
}

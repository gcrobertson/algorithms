package main

import (
	"fmt"
)

func main() {

	ex1()
	ex2()
	ex3()
}

func ex1() {

	input := 6
	output := true

	result := isUgly(input)

	fmt.Printf("Result expected? [%v]\n", output == result)
}

func ex2() {

	input := 8
	output := true

	result := isUgly(input)

	fmt.Printf("Result expected? [%v]\n", output == result)
}

func ex3() {

	input := 14
	output := false

	result := isUgly(input)

	fmt.Printf("Result expected? [%v]\n", output == result)
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Ugly Number.
// Memory Usage: 2.2 MB, less than 64.06% of Go online submissions for Ugly Number.
func isUgly(num int) bool {

	if num < 1 {
		return false
	}

	for _, factor := range []int{2, 3, 5} {
		for num%factor == 0 {
			num = num / factor
		}
	}

	return num == 1
}

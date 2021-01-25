package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	ex1()
	ex2()
	ex3()
}

func ex1() {
	input := 123
	expect := 321
	result := reverse(input)
	fmt.Printf("result [%v] matches expectation [%v]? [%v]\n", result, expect, result == expect)
}

func ex2() {
	input := -123
	expect := -321
	result := reverse(input)
	fmt.Printf("result [%v] matches expectation [%v]? [%v]\n", result, expect, result == expect)
}

func ex3() {
	input := 120
	expect := 21
	result := reverse(input)
	fmt.Printf("result [%v] matches expectation [%v]? [%v]\n", result, expect, result == expect)
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Integer.
// Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Reverse Integer.// Memory Usage: 2.2 MB
// O(n)
func reverse(x int) int {

	var result int

	// O(n)?
	for x != 0 {

		// O(1)
		pop := x % 10

		// O(1)
		x = x / 10

		// O(1)
		result = result*10 + pop

		// fmt.Printf("pop: %v\t x: %v\tresult: %v\n", pop, x, result)

		// O(1)
		if int(math.Abs(float64(result))) > math.MaxInt32 {
			return 0
		}
	}

	return result
}

// Runtime: 4 ms, faster than 40.59% of Go online submissions for Reverse Integer.
// Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Reverse Integer.
func reverseItoa(x int) int {

	//O(1)
	var negative bool
	//O(1)?
	s := strconv.Itoa(x)
	//O(1)
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}

	//O(1)
	var r string
	// O(n)
	for i := len(s) - 1; i > -1; i-- {
		r += string(s[i])
	}

	//O(1)?
	result, _ := strconv.Atoi(r)

	//O(1)
	if result > math.MaxInt32 {
		return 0
	}

	//O(1)
	if negative == true {
		//O(1)
		result *= -1
	}
	//O(1)
	return result

}

// Runtime: 4 ms, faster than 40.59% of Go online submissions for Reverse Integer.
// Memory Usage: 2.2 MB, less than 14.02% of Go online submissions for Reverse Integer.
// O(n)
func reverseStringBuilder(x int) int {

	// O(1)
	var negative bool
	// O(1)
	s := strconv.Itoa(x)
	// O(1)
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}

	// O(1)
	var r strings.Builder

	// O(n)
	for i := len(s) - 1; i > -1; i-- {
		fmt.Fprintf(&r, "%c", s[i])
	}

	// O(1)?
	result, _ := strconv.Atoi(r.String())

	// O(1)
	if result > math.MaxInt32 {
		return 0
	}

	// O(1)
	if negative == true {
		result *= -1
	}

	// O(1)
	return result
}

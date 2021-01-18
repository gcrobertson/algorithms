package main

import (
	"fmt"
	"math"
)

func main() {

	// "A"  -  1
	// "Z"  - 26
	// "AA" - 27
	// "AB" - 28
	// "AZ" - 56
	// "AC" - 57

	i := "A"
	e := 1
	r := titleToNumber(i)
	fmt.Printf("does result [%v] match expectation[%v]\n", r, r == e)

	i1 := "AB"
	e2 := 28
	r2 := titleToNumber(i1)
	fmt.Printf("does result [%v] match expectation[%v]\n", r2, r2 == e2)

	// (character value)

	// (character value) * 26 + (character value)

	i = "ZY"
	e = 701
	r = titleToNumber(i)
	fmt.Printf("does result [%v] match expectation[%v]\n", r, r == e)

	i = "AAA"
	e = 703
	r = titleToNumber(i)
	fmt.Printf("does result [%v] match expectation[%v]\n", r, r == e)

	//  y = character location in first character
	// x = character location in next character
	// f(x) = (26 * x[i]) + (26 * x[i+1]) ... + x[n]

	// ZY = 701

	// "AAA" (26 * 26) + 26 + 1		= min for the any string length
	// Output: 703

	// "AAAA" (26 * 26 * 26) + 26 + 1

	// "AA"

	// "AAA"
	// += 1 * 26
	// += 26 * 2
	// C = 3

}

func titleToNumber(s string) int {

	alphabet := make(map[string]int)

	// O(1) because it will always be constant regardless of len(s)
	for i := 1; i < 27; i++ {
		c := 64 + i // "A" = 65
		alphabet[string(c)] = i
	}

	var result int
	n := len(s)

	// O(n)
	for i := 0; i < n; i++ {
		curChar := string(s[n-1-i])
		result += alphabet[curChar] * int((math.Pow(float64(26), float64(i))))
	}
	return result
}

func titleToNumberR(s string) int {

	alphabet := make(map[string]int)

	// O(1) because it will always be constant regardless of len(s)
	for i := 1; i < 27; i++ {
		c := 64 + i // "A" = 65
		alphabet[string(c)] = i
	}

	var result int
	var exponent int

	// O(n)
	for i := len(s) - 1; i > -1; i-- {
		curChar := string(s[i])
		result += alphabet[curChar] * int((math.Pow(float64(26), float64(exponent))))
		exponent++
	}
	return result
}

// tuesday mongodb interview

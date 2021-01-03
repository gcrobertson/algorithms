package main

import (
	"fmt"
)

func main() {

	r1 := getSum(1, 2)
	r2 := getSum(-2, 3)

	fmt.Printf("did r1 [%v] return the correct solution? [%v]\n", r1, r1 == 3)
	fmt.Printf("did r2 [%v] return the correct solution? [%v]\n", r2, r2 == 1)
}

// See: [https://leetcode.com/problems/sum-of-two-integers/discuss/84323/Golang-concise-solution-with-an-explanation]
func getSum(a int, b int) int {

	/*	XOR [^]
	 *
	 *	RULE: if exclusively 1 of the input bits are 1, the resulting output is 1, otherwise the output is 0
	 *
	 *	APPLICATION: stands for the sum of each digit (ignoring the carry over)
	 */
	sum := a ^ b

	/*	AND [&]
	 *
	 *	RULE: if both input bits are 1, the resulting output is 1, otherwise the output is 0.
	 *
	 *	APPLICATION: stands for whether carry over occurs on each digit
	 */
	coefficient := a & b

	// carry over (a & b) << 1 repeatedly until no carry over occurred.
	for coefficient != 0 {
		coefficient = coefficient << 1
		t := sum
		sum = sum ^ coefficient
		coefficient = t & coefficient
	}

	return sum
}

func getSumTuple(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}

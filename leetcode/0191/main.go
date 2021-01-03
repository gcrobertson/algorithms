package main

import (
	"fmt"
	"strconv"
)

/*	`Hamming Weight`:	The Hamming weight of a string is the number of symbols that are different from the zero-symbol of the alphabet used.
 *
 *						For the most typical case, a string of bitss, this is the number of 1's in the string, or the digit sum
 *						of the binary representation of a given number.
 *
 *						Aliases:	`Population Count`
 *									`Popcount`
 *									`Sideways sum`
 *									`Bit summation`
 *
 *
 *	EXAMPLE:
 *	String			Hamming Weight
 *	11101			4
 *	11101000		4
 *  000000			0
 *  678012340567    10
 *
 *
 *	See: [https://en.wikipedia.org/wiki/Hamming_weight]
 */
func main() {

	ex1()
	// ex2()
}

func ex1() {
	// var input1 uint32
	input1 := uint32(00000000000000000000000000001011) //00000000000000000000000000001011
	result1 := hammingWeight(input1)
	fmt.Printf("Does result1 [%v] match expectation? [%v]\n", result1, result1 == 3)
}

func ex2() {
	var input2 uint32
	input2 = 00000000000000000000000010000000
	result2 := hammingWeight(input2)
	fmt.Printf("Does result2 [%v] match expectation? [%v]\n", result2, result2 == 1)
}

// Solution 1.
func hammingWeight(num uint32) int {

	s := strconv.FormatUint(uint64(num), 2)

	var result int

	for _, r := range s {
		// %c	the character represented by the corresponding Unicode code point
		if fmt.Sprintf("%c", r) == "1" {
			result++
		}
	}

	return result
}

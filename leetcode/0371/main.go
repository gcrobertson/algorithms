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

func getSum(a int, b int) int {

	for b != 0 {
		a, b = a^b, (a&b)<<1
	}

	return a
}

/*
func getSum(a, b int) int {



	binaryA := fmt.Sprintf("%08b", a)
	binaryB := fmt.Sprintf("%08b", b)

	carryOvers := ""
	runningSum := ""

	for i := 0; i < len(binaryA); i++ {

		// do all carry overs
		a, _ := strconv.Atoi(string(binaryA[i]))
		b, _ := strconv.Atoi(string(binaryB[i]))

		co := a & b
		if co == 1 {
			// we know to carry over ...
			carryOvers += "1"
		} else {
			carryOvers += "0"
		}

		// perform addition using the bitwise XOR operator
		if a^b == 1 {
			runningSum += "1"
		} else {
			runningSum += "0"
		}

	}

	fmt.Printf("a = [%v], b = [%v], carryover = [%v], runningsum = [%v]\n", a, b, carryOvers, runningSum)

	// fmt.Printf("binaryA was %v now %v\n", a, binaryA)
	// fmt.Printf("binaryB was %v now %v and is of type %T\n", a, binaryB, binaryB)

	return 0
}
*/

// Carry using the Bitwise AND &
// Bitwise math: Bitwise & only shows the operations where carry remainder is needed
// Carry is applied one position to the left of where it is discovered / created

// Exclusive OR, XOR ^
// Simulating addition can be done with the XOR operator.
//
//	1	0	1	0
//	1	1	0	0
//  =============
//	0	1	1	0
// XOR can perform the addition as seen above. The carry can be marked and a Bitwise shift can be used.

// Left shift operator <<
//
// 1010 << 1 == 0100
//
// The & operator can mark which positions need to carry a value.
// We apply the carry one position to the left of where it is discovered.
//
// In next iteration, we can apply these carry overs.

// & 	to find operations needing carry
// ^ 	to do addition operations in current iteration
// << 	to turn carry into what is applied in next interation

// variable A: hold running addition result
// variable B: hold carries

// Link: https://www.youtube.com/watch?v=qq64FrA2UXQ

// 1. find carries
// 2. do addition and store in A
// 3. b holds left-shifted carries

/*
 *		original:
 *			a = 1
 *			b = 3
 *
 *
 *		bitwise addition store carry overs with AND &:
 *			0001
 *			0011
 *			----
 *			0001
 *
 *		bitwise addition performed with XOR ^:
 *			0001
 *			0011
 *			----
 *			0010
 *
 *		left shift the carry result:
 *		0001 << 1 == 0010
 *
 *		bitwise addition store carry overs with AND &:
 *			0010
 *			0010
 *			----
 *			0010
 *
 *		bitwise addition performed with XOR ^:
 *			0010
 *
 *
 *
 *
 *
 *
 *
 *
 *
 */

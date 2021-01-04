package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	// i1 := 2
	// e1 := []int{0, 1, 1}
	// r1 := countBits(i1)

	// fmt.Printf("Does result [%v] match expecatation [%v]?\t [%v]\n", r1, e1, reflect.DeepEqual(r1, e1))

	i1 := 5
	e1 := []int{0, 1, 1, 2, 1, 2}
	r1 := countBits(i1)

	fmt.Printf("Does result [%v] match expecatation [%v]?\t [%v]\n", r1, e1, reflect.DeepEqual(r1, e1))

}

// O (n*sizeof(integer))
func countBitsMyAttempt(num int) []int {
	result := make([]int, num+1)
	for i := 0; i < num+1; i++ {
		binaryNumber := strconv.FormatInt(int64(i), 2)
		var c int
		for j := 0; j < len(binaryNumber); j++ {
			if string(binaryNumber[j]) == "1" {
				c++
			}
		}
		result[i] = c
	}
	return result
}

/// Approach #1 Pop Count
// O (n*k)
// for each integer `x` we need O(k) operations where `k` is the number of bits in `x`
func countBitsApproach1(num int) []int {
	result := make([]int, num+1)
	for i := 0; i < num+1; i++ {
		result[i] = approach1PopCount(i)
	}
	return result
}

func approach1PopCount(x int) int {
	var count int
	for count = 0; x != 0; count++ {

		/* @TODO: Circle back with Prakriti, not 100% sure why / how this zeroes out the least significant bit each time ...
		 *
		 *
		 */

		// fmt.Printf("x before = %v\t", x)
		x = x & (x - 1) // zeroing out the least significant nonzero bit
		// fmt.Printf("x after = %v\n", x)
	}

	return count
}

// Approach #2
func countBits(num int) []int {

	// See comment in : [https://leetcode.com/problems/counting-bits/solution/]
	// # num      0    1    2    3    4    5    6    7    8    9    10   11   12   13   14   15   16
	// # bin(num) 0000 0001 0010 0011 0100 0101 0110 0111 1000 1001 1010 1011 1100 1101 1110 1111 10000
	// # count    0    1    1    2    1    2    2    3    1    2    2    3    2    3    3    4    1
	// # pattern  0    1    0+1  1+1  0+1  1+1  1+1  2+1  0+1  1+1  1+1  2+1  1+1  2+1  2+1  3+1  0+1
	// # i                  2         4                   8                                       16
	// # j        0    1    0    1    0    1    2    3    0    1    2    3    4    5    6    7    0

	result := make([]int, num+1)

	i := 0
	b := 1

	// [0,b)
	for b <= num {

		// generate [b, 2b) or [b, num) from [0, b)
		for i < b && i+b <= num {

			result[i+b] = result[i] + 1
			i++
		}

		i = 0   // reset i
		b <<= 1 // b = 2b
	}

	return result
}

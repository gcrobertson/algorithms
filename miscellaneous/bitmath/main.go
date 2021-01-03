package main

import "fmt"

func main() {

	x := 5

	fmt.Printf("the value of x is: %b\n", x)

	// bitwiseAND(x)
	// bitwiseOR(x)
	// bitwiseXOR(x)
	bitwiseNOT(x)
}

// If both input bits are 1, the resulting output is 1, otherwise the output is 0.
func bitwiseAND(x int) {
	fmt.Printf("\nPerforming bitwise AND [ & ]\n")

	for i := 0; i < 6; i++ {
		y := x & i
		fmt.Printf("the value of i: %v y in base 10: %d y in binary: %b\n", i, y, y)
	}

	// i		x		y
	// 0		101		0
	// 1		101		1
	// 10		101		0
	// 11		101		1
	// 100		101		100
	// 101		101		101
}

// If either or both input bits are 1, the resulting output is 1, otherwise the output is 0.
func bitwiseOR(x int) {
	fmt.Printf("\nPerforming bitwise OR [ | ]\n")

	for i := 0; i < 6; i++ {
		y := x | i
		fmt.Printf("the values of i: %b \ty: %b\n", i, y)
	}

	// i		x		y
	// 0		101		101
	// 1		101		101
	// 10		101		111
	// 11		101		111
	// 100		101		101
	// 101		101		101

}

// If exclusively 1 of the input bits are 1, the resulting output is 1, otherwise the output is 0.
func bitwiseXOR(x int) {
	fmt.Printf("\nPerforming bitwise XOR [ ^ ]\n")

	for i := 0; i < 6; i++ {
		y := x ^ i
		fmt.Printf("the values of i: %b \ty: %b\n", i, y)
	}

	// i		x		y
	// 0		101		101
	// 1		101		100
	// 10		101		111
	// 11		101		110
	// 100		101		001
	// 101		101		000

}

// Bitwise NOT changes each bit to its opposite.
func bitwiseNOT(x int) {
	fmt.Printf("\nPerforming bitwise NOT [ ~ ]\n")

	var y byte

	y &^= byte(5)

	fmt.Printf("%08b\t%08b\n", byte(5), y)

	// for i := 0; i < 6; i++ {
	// 	var y int
	// 	y = ^i

	// 	fmt.Printf("the values of i: %b \ty: %b\n", i, y)
	// }

	// i		y
	// 0		1
	// 1		0
	// 10		1
	// 11		0
	// 100		11
	// 101		10
}

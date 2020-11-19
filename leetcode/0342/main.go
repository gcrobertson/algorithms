package main

import "fmt"

func main() {

	fmt.Printf("Is 16 a power of four? %v\n", isPowerOfFour(16))
	fmt.Printf("Is 26 a power of four? %v\n", isPowerOfFour(26))
}

func isPowerOfFour(n int) bool {

	// 4 ^ 0 = 1
	// 4 ^ 1 = 4
	// 4 ^ 2 = 16
	// 4 ^ 3 = 64
	// 4 ^ 4 = 256
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	for n != 1 {

		if n%4 != 0 {
			return false
		}
		n = n / 4
	}

	return true
}

package main

import (
	"fmt"
)

func main() {

	var input, expect uint32

	input = 43261596
	expect = 964176192

	r := reverseBits(input)

	fmt.Printf("Did the result [%v] match the expectation? [%v]\n", r, expect == r)
}

func reverseBits(num uint32) uint32 {

	// fmt.Printf("%032b\n", num)

	var rev uint32

	for i := 0; i < 32; i++ {

		rev = rev<<1 + num&1
		num >>= 1
	}

	return rev
}

// REVISIT this one with Prakriti!

package main

import (
	"fmt"
)

func main() {

	var n uint32

	// n = 00000010100101000001111010011100

	n = 7 // 111

	r := reverseBits(n)

	fmt.Printf("Did the result [%v] match the expectation? [%v]\n", r, n == r)
}

func reverseBits(num uint32) uint32 {

}

func reverseBitsAnswerProvided(num uint32) uint32 {

	var rev uint32

	for i := 0; i < 32; i++ {

		rev = rev<<1 + num&1
		num >>= 1
	}

	return rev
}

// s := strconv.FormatUint(uint64(num), 2)

// // convert
// bits := strconv.FormatUint(uint64(num), 2)

// fmt.Println(bits)

/*
	var bit byte = 0x0F
	fmt.Printf("%08b\n", bit)
	fmt.Printf("%08b\n", ^bit)
*/

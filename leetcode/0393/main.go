package main

import (
	"fmt"
	"strconv"
)

func main() {
	// UTF-8 can be anywhere from 1 byte to 4 bytes
	//
	// we are going to be given this data that is: 1-4 bytes long
	// example 1:
	//  00000001                            = 1 byte
	//  11000000 10000011                   = 2 bytes
	//  11100000 10000011 10000011          = 3 bytes
	//  11110000 10000011 10000011 10000011 = 4 bytes

	data1 := []int{197, 130, 1}
	data2 := []int{235, 140, 4}

	fmt.Printf("The result [%v] should be true\n", validUtf8(data1))

	fmt.Printf("The result [%v] should be false\n", validUtf8(data2))
}

// Written by Gregory Robertson and Prakriti Khanal
// Let's go Palantir!
func validUtf8(data []int) bool {

	var noOfBitsToProcess int
	for _, v := range data {

		bits := strconv.FormatInt(int64(v), 2)

		if len(bits) > 8 {
			index := len(bits) - 8
			bits = bits[index:]
		} else {
			bits = fmt.Sprintf("%08v\n", bits)
		}

		if noOfBitsToProcess == 0 {

			// calculate the bits to process
			for i := 0; i < len(bits); i++ {

				if string(bits[i]) == "0" {
					break
				}
				noOfBitsToProcess++
			}

			if noOfBitsToProcess > 4 || noOfBitsToProcess == 1 {
				return false
			}

			if noOfBitsToProcess == 0 {
				continue
			}

		} else {

			if string(bits[0]) != "1" || string(bits[1]) != "0" {
				return false
			}

		}

		noOfBitsToProcess--
	}

	return noOfBitsToProcess == 0

}

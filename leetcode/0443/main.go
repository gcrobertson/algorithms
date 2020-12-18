package main

import "fmt"

func main() {

	chars := []byte{97, 97, 98, 98, 99, 99, 99} //aabbccc			//a2b2c3

	result := compress(chars)

	expect := 6

	fmt.Printf("Does the result [%v] match the expectation [%v]? [%v]\n", result, expect, result == expect)

	chars = chars[0:result]

	fmt.Println(string(chars))

}

func compress(chars []byte) int {
	var writer, reader int
	for reader < len(chars) {
		var count int
		chars[writer] = chars[reader]
		for reader < len(chars) && chars[reader] == chars[writer] {
			reader++
			count++
		}
		// no longer the same character
		if count > 1 {
			number := fmt.Sprintf("%d", count)
			for _, digit := range []byte(number) {
				writer++
				chars[writer] = digit
			}
		}
		writer++
	}
	return writer
}

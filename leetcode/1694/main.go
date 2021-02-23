package main

import (
	"fmt"
	"strings"
)

func main() {

	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
}

func ex1() {
	input := "1-23-45 6"
	expect := "123-456"
	result := reformatNumber(input)

	fmt.Printf("Result [%v] matches expectation [%v]? [%v]\n", result, expect, result == expect)
}

func ex2() {
	input := "123 4-567"
	expect := "123-45-67"
	result := reformatNumber(input)

	fmt.Printf("Result [%v] matches expectation [%v]? [%v]\n", result, expect, result == expect)
}

func ex3() {
	input := "123 4-5678"
	expect := "123-456-78"
	result := reformatNumber(input)

	fmt.Printf("Result [%v] matches expectation [%v]? [%v]\n", result, expect, result == expect)
}

func ex4() {
	input := "12"
	expect := "12"
	result := reformatNumber(input)

	fmt.Printf("Result [%v] matches expectation [%v]? [%v]\n", result, expect, result == expect)
}

func ex5() {
	input := "--17-5 229 35-39475 "
	expect := "175-229-353-94-75"
	result := reformatNumber(input)

	fmt.Printf("Result [%v] matches expectation [%v]? [%v]\n", result, expect, result == expect)
}

// Runtime: 4 ms, faster than 100.00% of Go online submissions for Reformat Phone Number.
// Memory Usage: 2.2 MB, less than 39.85% of Go online submissions for Reformat Phone Number. without the string.Builder

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reformat Phone Number.
// Memory Usage: 2.1 MB, less than 79.71% of Go online submissions for Reformat Phone Number. with the string.Builder
func reformatNumber(number string) string {

	number = strings.ReplaceAll(number, " ", "")
	number = strings.ReplaceAll(number, "-", "")

	var result strings.Builder
	var index int

	last := len(number) - 1

	for index < len(number) {

		// if there are at least 4 characters left
		if last-index > 3 {

			fmt.Fprintf(&result, "%s", number[index:index+3]+"-")
			index += 3

		} else {

			// finalize what is left
			if last-index == 3 {

				fmt.Fprintf(&result, "%s", number[index:index+2]+"-"+number[index+2:])

				return result.String()

			}

			fmt.Fprintf(&result, "%s", number[index:])

			return result.String()
		}

	}

	return result.String()
}

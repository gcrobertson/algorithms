package main

import "fmt"

func ex1() {
	input := 234
	expect := 15
	result := subtractProductAndSum(input)
	fmt.Printf("Example 1? %v\n", expect == result)
}

func ex2() {
	input := 4421
	expect := 21
	result := subtractProductAndSum(input)
	fmt.Printf("Example 2? %v\n", expect == result)
}

func main() {
	ex1()
	ex2()
}

func subtractProductAndSum(n int) int {

	// uses modulus operator

	product := 1
	var sum int

	for n > 0 {

		product *= n % 10
		sum += n % 10

		// fmt.Printf("product = %v\n", product)
		// fmt.Printf("sum: %v\n", sum)

		n /= 10
	}

	return product - sum
}

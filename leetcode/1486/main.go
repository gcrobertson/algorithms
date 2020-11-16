package main

import "fmt"

func main() {

	n := 5
	start := 0
	output := 8

	r := xorAdd(n, start)

	fmt.Printf("are these the same? %v %v\n", output, r)

	fmt.Printf("%v %v", 60, 60<<2)
}

func xorAdd(input, start int) int {

	var val int

	for i := 0; i < input; i++ {
		val ^= start + 2*i
	}

	return val
}

// This is not the solution but it is XORADD operations in Golang, a first step to the solution maybe.

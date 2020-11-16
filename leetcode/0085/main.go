package main

import "fmt"

/*
	Given a 2D binary matrix filled with O & 1, find the largest rectangle containing only 1s and return its area.

	Example:

	Input:
	[
		[1,0,1,0,0]
		[1,0,1,1,1]
		[1,1,1,1,1]
		[1,0,0,1,0]
	]
	Output: 6

*/

func main() {

	input := [][]string{
		[]string{"1", "0", "1", "0", "0"},
		[]string{"1", "0", "1", "1", "1"},
		[]string{"1", "1", "1", "1", "1"},
		[]string{"1", "0", "0", "1", "0"},
	}

	for _, v := range input {
		fmt.Println(v)
	}

	output := maximalRectangle(input)

	fmt.Printf("Area is: %v\n", output)
}

func maximalRectangle(input [][]string) int {

	return 0
}

package main

import "fmt"

// 	https://yourbasic.org/golang/implement-stack/

/*	Basic stack [LIFO] data structure
 *
 *	Push:	append()
 *	Pop:	slice the slice
 */
func main() {

	var stack []string

	stack = append(stack, "world!") // Push
	stack = append(stack, "Hello ") // Push

	// Iterate descending order [last in first out]
	for len(stack) > 0 {

		n := len(stack) - 1 // Last element pushed to stack
		fmt.Print(stack[n])

		stack[n] = ""     // Erase element to avoid memory leaks
		stack = stack[:n] // Pop
	}
}

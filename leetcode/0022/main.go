package main

import (
	"fmt"
	"reflect"
)

func main() {
	ex1()
	ex2()
}

func ex1() {

	input := 3
	expect := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	result := generateParenthesis(input)
	fmt.Printf("Does result %v match expectation? [%v]\n", result, reflect.DeepEqual(expect, result))
}

func ex2() {

	input := 1
	expect := []string{"()"}
	result := generateParenthesis(input)
	fmt.Printf("Does result %v match expectation? [%v]\n", result, reflect.DeepEqual(expect, result))
}

// https://leetcode.com/problems/generate-parentheses/discuss/807298/Golang-4ms-3.4MB-easy-solution
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	result := gen("(", 1, 0, n)
	return result
}

func gen(prefix string, open int, closed int, n int) []string {
	if closed == n {
		return []string{prefix}
	}

	var left []string
	if open < n {
		left = gen(prefix+"(", open+1, closed, n)
	}

	var right []string
	if closed < open {
		right = gen(prefix+")", open, closed+1, n)
	}

	return append(left, right...)
}

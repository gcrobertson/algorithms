package main

import (
	"fmt"
	"math"
)

func main() {

	ex1()
	ex2()
	ex3()
	ex4()
}

func ex1() {
	input := 0
	expect := 0
	result := fib(input)
	fmt.Printf("input [%v] result [%v] matches expectation? [%v]\n", input, result, expect == result)
}

func ex2() {
	input := 1
	expect := 1
	result := fib(input)
	fmt.Printf("input [%v] result [%v] matches expectation? [%v]\n", input, result, expect == result)
}

func ex3() {
	input := 3
	expect := 2
	result := fib(input)
	fmt.Printf("input [%v] result [%v] matches expectation? [%v]\n", input, result, expect == result)
}

func ex4() {
	input := 4
	expect := 3
	result := fib(input)
	fmt.Printf("input [%v] result [%v] matches expectation? [%v]\n", input, result, expect == result)
}

// F(0) = 0, F(1) = 1
// F(n) = F(n - 1) + F(n - 2), for n > 1.
// 	0	1	2	3	4	5	6
//	0	1	1	2	3	5	8

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Fibonacci Number.
// Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Fibonacci Number.
// O(n)
func fibON(n int) int {

	if n < 2 {
		return n
	}

	s1 := 0
	s2 := 1

	for i := 2; i < n; i++ {
		t1 := s1
		t2 := s2
		s2 = t1 + t2
		s1 = t2
	}

	return s1 + s2
}

// Recursion
// O(2ⁿ)
// Runtime: 12 ms, faster than 19.00% of Go online submissions for Fibonacci Number.
// Memory Usage: 1.9 MB, less than 26.16% of Go online submissions for Fibonacci Number.
func fibRecursion(n int) int {

	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

// O(n) if unvisited, O(1) when visited and stored later on ...
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Fibonacci Number.
// Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Fibonacci Number.
func fibCache(n int) int {

	if n < 2 {
		return n
	}

	return memoization(n)
}

var cache = make(map[int]int, 30)

func memoization(n int) int {

	if len(cache) == 0 {

		n1 := 0
		n2 := 1

		cache[0] = n1
		cache[1] = n2

		// 0 <= n <= 30
		for i := 2; i < 31; i++ {
			cache[i] = cache[i-1] + cache[i-2]
		}
	}

	return cache[n]
}

/*	golden ratio aka Binet's formula:
 *	φ = (1 + √5) / 2 ≈ 1.161803
 *
 *	fibonacci number for `n`, f(n) = φ ^ (n/√5)
 *
 */
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Fibonacci Number.
// Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Fibonacci Number.
// O(1) but not quite, cannot perform Xⁿ in constant time.
func fib(n int) int {

	var goldenRatio float64 = float64((1 + math.Sqrt(5)) / 2)
	var exponent float64 = float64(n)

	return int(math.Round(math.Pow(goldenRatio, exponent) / math.Sqrt(5)))
}

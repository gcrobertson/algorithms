package main

import (
	"fmt"
	"strconv"
)

func main() {

	ex1()
	ex2()
	ex3()
}

func ex1() {
	ops := []string{"5", "2", "C", "D", "+"}
	expect := 30
	result := calPoints(ops)
	fmt.Printf("Does the result [%v] match the expectations? [%v]\n", result, result == expect)
}

func ex2() {
	ops := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	expect := 27
	result := calPoints(ops)
	fmt.Printf("Does the result [%v] match the expectations? [%v]\n", result, result == expect)
}

func ex3() {
	ops := []string{"1"}
	expect := 1
	result := calPoints(ops)
	fmt.Printf("Does the result [%v] match the expectations? [%v]\n", result, result == expect)
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Baseball Game.
// Memory Usage: 2.6 MB, less than 33.33% of Go online submissions for Baseball Game.
func calPoints(ops []string) int {

	var ints []int

	for _, op := range ops {

		switch op {

		case "+": // "+": Record a new score that is the sum of the previous two scores.
			i1 := ints[len(ints)-1]
			i2 := ints[len(ints)-2]
			ints = append(ints, i1+i2)

		case "D": // "D": Record a new score that is double the previous score.

			i1 := ints[len(ints)-1] * 2
			ints = append(ints, i1)

		case "C": // "C": Invalidate the previous score, removing it from the record.
			ints = ints[0 : len(ints)-1]

		default:
			i, _ := strconv.Atoi(op)
			ints = append(ints, i)
		}
	}

	var sum int

	for _, i := range ints {
		sum += i
	}

	// 	Return the sum of all the scores on the record
	return sum
}

package main

import (
	"fmt"
	"math"
)

func main() {
	// inp := []int{2, 5}
	input := []int{4, 8}
	// intervals := [][]int{
	// 	// []int{1, 3},
	// 	// []int{6, 9},
	// 	[]int{1, 2},
	// 	[]int{3, 5},
	// 	[]int{6, 7},
	// 	[]int{8, 10},
	// 	[]int{12, 16},
	// }
	// result := insertInterval(intervals, input)
	// fmt.Printf("the result of adding non-overlapping interval %v to %v:\n%v\n", input, intervals, result)

	r := [][]int{
		[]int{1, 2},
		[]int{3, 5},
		[]int{6, 7},
		[]int{8, 10},
		[]int{12, 16},
	}

	fmt.Printf("\nbefore adding %v: %+v\n", input, r)
	insertIntoSlice(r, input)
	fmt.Printf("\n\nafter adding %v: %+v\n", input, r)
}

func insertInterval(intervals [][]int, inp []int) [][]int {

	var index int
	var r [][]int

	for _, interval := range intervals {
		if inp[0] > interval[1] {
			// precedes
			r = append(r, interval)
			index++
		}
	}

	var max, min int
	for index < len(intervals) && intervals[index][0] <= inp[1] {
		min = int(math.Min(float64(intervals[index][0]), float64(inp[0])))
		max = int(math.Max(float64(intervals[index][1]), float64(inp[1])))
		index++
	}
	r = append(r, []int{min, max})

	for index < len(intervals) {
		r = append(r, intervals[index], inp)
		index++
	}

	return r
}

func insertIntoSlice(intervals [][]int, input []int) {

	// let us try to do the same thing but with the same slice.

	var index int

	for index < len(intervals) && input[0] > intervals[index][1] {
		// precedes
		index++
	}

	insert := index
	var min, max int
	for {

		if index >= len(intervals) {
			break
		}

		if input[1] >= intervals[index][0] {

			min = int(math.Min(float64(intervals[index][0]), float64(input[0])))
			max = int(math.Max(float64(intervals[index][1]), float64(input[1])))

			fmt.Printf("before deletion index is %v and len is %v\n", index, len(intervals))

			copy(intervals[index:], intervals[index+1:])
			intervals = intervals[:len(intervals)-1]

			fmt.Printf("after deletion len is: %v and the slice is %+v\n", len(intervals), intervals)

			index++
		}
	}

	intervals[insert] = []int{min, max}
}

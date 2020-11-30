package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	input := [][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}

	output := merge(input)

	fmt.Printf("The output is %v\n", output)
}

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var merged [][]int
	merged = append(merged, intervals[0])

	fmt.Printf("merged: %v\n", merged)

	for _, interval := range intervals {

		starti := interval[0]
		endi := interval[1]

		last := merged[len(merged)-1][1]

		if starti > last {

			merged = append(merged, interval)

		} else {

			for k, v := range merged {

				startm := v[0]
				endm := v[1]

				if starti >= startm && starti <= endm {
					merged[k] = []int{startm, int(math.Max(float64(endm), float64(endi)))}
				}
			}
		}

	}

	return merged
}

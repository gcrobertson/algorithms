package main

import (
	"math"
	"sort"
)

func partitionLabelsAttempt1(S string) []int {

	// map ["c"][start int, end int]
	cmap := make(map[string][]int, 26)

	// O(n)
	for key, val := range S {

		if _, ok := cmap[string(val)]; !ok {
			cmap[string(val)] = []int{key, key}
		} else {
			cmap[string(val)][1] = key
		}
	}

	//fmt.Println(cmap)

	// O(k) where k is # of unique characters in given string, S
	var mergeslice [][]int
	for _, interval := range cmap {
		mergeslice = append(mergeslice, interval)
	}

	// O(log 2 k) Sort the merge slice I think?
	sort.Slice(mergeslice, func(i, j int) bool { return mergeslice[i][0] < mergeslice[j][0] })

	//fmt.Println("merge slice", mergeslice)

	// perform the merge
	var merged [][]int
	merged = append(merged, mergeslice[0])

	for _, interval := range mergeslice {

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

	//fmt.Printf("merged: %v\n", merged)

	var result []int

	for _, interval := range merged {

		add := interval[1] + 1 - interval[0]
		result = append(result, add)
	}

	return result
}

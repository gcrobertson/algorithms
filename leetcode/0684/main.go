package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	// input := [][]int{
	// 	[]int{1, 2},
	// 	[]int{1, 3},
	// 	[]int{2, 3},
	// }

	input := [][]int{
		[]int{1, 2},
		[]int{2, 3},
		[]int{3, 4},
		[]int{1, 4},
		[]int{1, 5},
	}

	r := findRedundantConnection(input)

	fmt.Printf("redundant connection is: %v", r)
}

func find(set []int, x int) int {
	if set[x] == x {
		return x
	}
	return find(set, set[x])
}

func findRedundantConnection(edges [][]int) []int {
	set := make([]int, len(edges)+1)
	for i := range set {
		set[i] = i
	}
	for _, edge := range edges {
		px := find(set, edge[0])
		py := find(set, edge[1])
		if px == py {
			return edge
		}

		// union
		set[px] = py
	}
	return []int{}
}

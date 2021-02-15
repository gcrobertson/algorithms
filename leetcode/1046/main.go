package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {

	// i := []int{2, 7, 4, 1, 8, 1}
	i := []int{2, 2}
	e := 1
	r := lastStoneWeight(i)

	fmt.Printf("expected[%v]? [%v]", r, e == r)
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Last Stone Weight.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Last Stone Weight.
func lastStoneWeightSlice(stones []int) int {

	if len(stones) == 0 {
		return 0
	}

	// O(log n)
	sort.Ints(stones)

	// O(n)
	for len(stones) > 1 {

		x, y := stones[len(stones)-2], stones[len(stones)-1]

		if x == y {
			stones = stones[0 : len(stones)-2]
		} else {
			stones[len(stones)-2] = y - x
			stones = stones[0 : len(stones)-1]
			sort.Ints(stones) // O(log n)
		}
	}

	if len(stones) == 0 {
		return 0
	}

	return stones[0]
}

/*	https://pkg.go.dev/container/heap
 *
 *
 *
 *
 *
 *
 *
 *
 */

// intHeap ...
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] } // ?
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	// push & pop use pointer receivers because they modify the slice length, not just its content

	// type assertion: https://golang.org/ref/spec#Type_assertions
	// i, _ := x.(int)

	// *h = append(*h, i)
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {

	pq := intHeap(stones)

	heap.Init(&pq)

	for pq.Len() > 1 {
		pq.Push(&pq)
	}

	return heap.Pop(&pq).(int)
}

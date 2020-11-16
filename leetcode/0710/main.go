package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(N, blacklist);
 * param_1 := obj.Pick();
 */
func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	obj := Constructor(5, []int{2, 5})
	param1 := obj.Pick()
	fmt.Printf("Randomly selected %v\n", param1)

}

// Solution ...
type Solution struct {
	n int
	b []int
}

// Constructor ...
func Constructor(N int, blacklist []int) Solution {

	b := sort.IntSlice(blacklist)
	b.Sort()

	return Solution{
		n: N,
		b: b,
	}

}

// Pick ...
func (sol *Solution) Pick() int {

	random := 5 // rand.Intn(sol.n - len(sol.b))

	var lo int
	hi := len(sol.b) - 1

	for lo < hi {

		mid := (lo + hi + 1) / 2

		// the number of whitelist numbers less than sol.b[mid]
		wl := sol.b[mid] - mid

		if wl > random {
			hi = mid - 1
		} else {
			lo = mid
		}
	}

	if lo == hi && sol.b[lo]-lo <= random {
		return random + lo + 1
	}

	return random
}

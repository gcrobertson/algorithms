package main

import (
	"fmt"
	"reflect"
)

func main() {

	s := "leetcode"
	e := 0
	r := firstUniqChar(s)

	fmt.Printf("expectation matches result [%v] ? [%v]\n", r, reflect.DeepEqual(e, r))

	s = "loveleetcode"
	e = 2
	r = firstUniqChar(s)

	fmt.Printf("expectation matches result [%v] ? [%v]\n", r, reflect.DeepEqual(e, r))
}

// Runtime: 28 ms, faster than 66.99% of Go online submissions for First Unique Character in a String.
// Memory Usage: 5.8 MB, less than 53.85% of Go online submissions for First Unique Character in a String.
// O(n+n) with space O(k) where k is unique characters in n?
func firstUniqChar(s string) int {
	// O(1)
	hash := make(map[string]int)

	// O(n)
	for _, char := range s {
		hash[string(char)]++
	}

	// O(n)
	for idx, char := range s {
		if hash[string(char)] == 1 {
			return idx
		}
	}

	return -1
}

// Runtime: 4 ms, faster than 100.00% of Go online submissions for First Unique Character in a String.
// Memory Usage: 5.8 MB, less than 53.85% of Go online submissions for First Unique Character in a String.
func firstUniqCharEvenFaster(s string) int {

	// O(k) where k == 26
	flags := make([]int, 26)
	for i := range flags {
		flags[i] = -1
	}

	// O(n)
	slen := len(s)
	for i, ch := range s {
		idx := byte(ch - 'a')
		if flags[idx] == -1 {
			flags[idx] = i
		} else {
			flags[idx] = slen
		}
	}

	// O(k), this works because it is a slice in order and not a normal map data structure
	min := slen
	for i := range flags {
		if flags[i] != -1 && flags[i] < min {
			min = flags[i]
		}
	}

	if min == slen {
		return -1
	}
	return min
}

package main

/*
Given:
	- string S
	- string T

Find:
 - the minimum window in S in which will contain all the characters in T

Solve in:
	- O(n)

NOTE:
	- if no such window in S, return ""
 	- if such window there will always be only 1 unique minimum window in S
*/

import (
	"fmt"
)

type trieNode struct {
	value    string // "A"
	children [26]*trieNode
}

func main() {

	s := "ABODECODEBANC"
	t := "ABC"
	output := "BANC"

	r := minWindowBruteForce(s, t)
	fmt.Printf("was response %v correct? %v\n", r, r == output)
	// r := minWindow(s, t)
	// fmt.Printf("was response %v correct? %v\n", r, r == output)

}

func minWindow(s, t string) string {
	result := s
	var found bool
	var lptr, rptr int

	for lptr < len(s) && rptr < len(s) {
		test := s[lptr : rptr+1]
		resp := testMap(test, t)
		if resp == false {
			if rptr+1 == len(s) {
				lptr++
				rptr = lptr
			} else {
				rptr++
			}
		} else {
			found = true
			if len(test) < len(result) {
				result = test
			}
			lptr++
		}
	}

	if found == false {
		return ""
	}

	return result
}

func testMap(s, t string) bool {
	hash := respMap(t)
	for _, v := range s {
		if _, ok := hash[string(v)]; ok {
			hash[string(v)]--
		}
	}
	for _, v := range hash {
		if v > 0 {
			return false
		}
	}
	return true
}

func respMap(t string) map[string]int {
	hash := make(map[string]int)
	for _, v := range t {
		hash[string(v)]++
	}
	return hash
}

func respMapSolved(hash map[string]int) bool {
	for _, v := range hash {
		if v > 0 {
			return false
		}
	}
	return true
}

func minWindowBruteForce(s, t string) string {
	var result string
	var windows []string

	for l := 0; l < len(s); l++ {
		window := string(s[l])
		for r := l + 1; r < len(s); r++ {
			window += string(s[r])
		}
		windows = append(windows, string(window))
	}

	for _, window := range windows {

		nhash := respMap(t)

		fmt.Printf("\nwindow: %v\n", window)
		for _, v := range window {
			if _, ok := nhash[string(v)]; ok {
				nhash[string(v)]--
			}
		}

		if respMapSolved(nhash) {
			fmt.Printf("%v is a solution with length %v\n", window, len(window))
			result = window
		}
	}

	return result

	// 	window: ABODECODEBANC
	// ABODECODEBANC is a solution with length 13

	// window: BODECODEBANC
	// BODECODEBANC is a solution with length 12

	// window: ODECODEBANC
	// ODECODEBANC is a solution with length 11

	// window: DECODEBANC
	// DECODEBANC is a solution with length 10

	// window: ECODEBANC
	// ECODEBANC is a solution with length 9

	// window: CODEBANC
	// CODEBANC is a solution with length 8

	// window: ODEBANC
	// ODEBANC is a solution with length 7

	// window: DEBANC
	// DEBANC is a solution with length 6

	// window: EBANC
	// EBANC is a solution with length 5

	// window: BANC
	// BANC is a solution with length 4

	// window: ANC

	// window: NC

	// window: C
	// was response BANC correct? true

}

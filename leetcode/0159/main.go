package main

import "fmt"

func main() {
	ex1()
	ex2()
}

// Given a string s, return the length of the longest substring that contains at most 2 distinct characters.

func ex1() {

	s := "eceba"
	e := 3
	r := lengthOfLongestSubstringTwoDistinct(s)
	fmt.Printf("%v\n", e == r)

}

func ex2() {

	s := "ccaabbb"
	e := 5
	r := lengthOfLongestSubstringTwoDistinct(s)
	fmt.Printf("%v\n", e == r)

}

// 1 <= s.length <= 104
// s consists of English letters.
func lengthOfLongestSubstringTwoDistinct(s string) int {

	var p1, p2 int
	var max int

	hash := make(map[byte]int)

	for p2 < len(s) {

		c := s[p2]
		hash[c]++
		p2++

		for len(hash) > 2 {

			if p2-1-p1 > max {
				max = p2 - 1 - p1
			}

			hash[s[p1]]--

			if hash[s[p1]] == 0 {
				delete(hash, s[p1])
			}

			p1++
		}

		if p2-p1 > max {
			max = p2 - p1
		}
	}

	return max
}

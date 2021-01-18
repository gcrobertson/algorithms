package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "leetcodeisacommunityforcoders"
	e := "ltcdscmmntyfrcdrs"
	r := removeVowelsSB(s)

	fmt.Printf("result matches expectation? [%v]\n", e == r)

	s = "aeiou"
	e = ""
	r = removeVowelsSB(s)

	fmt.Printf("result matches expectation? [%v]\n", e == r)
}

var vowels = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Vowels from a String.
// Memory Usage: 2.9 MB, less than 32.14% of Go online submissions for Remove Vowels from a String.
func removeVowels(s string) string {

	var r string
	// o(n)
	for _, c := range s {
		// O(1)
		if _, vowel := vowels[c]; !vowel {
			// ? use a string builder
			r += string(c)
		}
	}

	return r
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Vowels from a String.
Memory Usage: 2.1 MB, less than 50.00% of Go online submissions for Remove Vowels from a String.
*/
func removeVowelsSB(s string) string {

	var r strings.Builder
	// o(n)
	for _, c := range s {
		// O(1)
		if _, ok := vowels[c]; !ok {
			// use a string builder
			fmt.Fprintf(&r, "%s", string(c))
		}
	}

	return r.String()
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Vowels from a String.
// Memory Usage: 2.1 MB, less than 42.86% of Go online submissions for Remove Vowels from a String.
func removeVowelsSwitch(s string) string {

	var r strings.Builder
	// o(n)
	for _, c := range s {
		// O(1)
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			break
		default:
			// use a string builder
			fmt.Fprintf(&r, "%s", string(c))
		}
	}

	return r.String()
}

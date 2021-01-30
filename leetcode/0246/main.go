package main

import "fmt"

func main() {

	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
}

func ex1() {
	input := "69"
	expect := true
	output := isStrobogrammatic(input)
	fmt.Printf("does result [%v] match expectation? [%v]\n", output, output == expect)
}

func ex2() {
	input := "88"
	expect := true
	output := isStrobogrammatic(input)
	fmt.Printf("does result [%v] match expectation? [%v]\n", output, output == expect)
}

func ex3() {
	input := "962"
	expect := false
	output := isStrobogrammatic(input)
	fmt.Printf("does result [%v] match expectation? [%v]\n", output, output == expect)
}

func ex4() {
	input := "1"
	expect := true
	output := isStrobogrammatic(input)
	fmt.Printf("does result [%v] match expectation? [%v]\n", output, output == expect)
}

func ex5() {
	input := "2"
	expect := false
	output := isStrobogrammatic(input)
	fmt.Printf("does result [%v] match expectation? [%v]\n", output, output == expect)
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Strobogrammatic Number.
// Memory Usage: 1.9 MB, less than 25.00% of Go online submissions for Strobogrammatic Number.
// var ref = map[string]string{
// 	"0": "0",
// 	"1": "1",
// 	"6": "9",
// 	"8": "8",
// 	"9": "6",
// }

func isStrobogrammatic(num string) bool {

	// Runtime: 0 ms, faster than 100.00% of Go online submissions for Strobogrammatic Number.
	// Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Strobogrammatic Number.
	var ref = map[string]string{
		"0": "0",
		"1": "1",
		"6": "9",
		"8": "8",
		"9": "6",
	}

	if len(num) == 0 {
		return false
	}

	// can solve with 2 pointers
	p1, p2 := 0, len(num)-1

	// O(n/2)
	for p1 <= p2 {

		i := string(num[p1])             // index for ref
		v, ok := ref[i]                  // v = strobogrammatic value
		if !ok || v != string(num[p2]) { // if there is no reference key, e.g. "5", reference != given string key, return false
			return false
		}

		p1++
		p2--
	}

	return true
}

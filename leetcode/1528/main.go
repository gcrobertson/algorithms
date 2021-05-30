package main

import "fmt"

func ex1() {
	input := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	expect := "leetcode"
	result := restoreString(input, indices)
	fmt.Printf("Example 1: %v\n", result == expect)
}

func ex2() {
	input := "abc"
	indices := []int{0, 1, 2}
	expect := "abc"
	result := restoreString(input, indices)
	fmt.Printf("Example 1: %v\n", result == expect)
}

func ex3() {
	input := "aiohn"
	indices := []int{3, 1, 4, 2, 0}
	expect := "nihao"
	result := restoreString(input, indices)
	fmt.Printf("Example 1: %v\n", result == expect)
}

func ex4() {
	input := "aaiougrt"
	indices := []int{4, 0, 2, 6, 7, 3, 1, 5}
	expect := "arigatou"
	result := restoreString(input, indices)
	fmt.Printf("Example 1: %v\n", result == expect)
}

func ex5() {
	input := "art"
	indices := []int{1, 0, 2}
	expect := "rat"
	result := restoreString(input, indices)
	fmt.Printf("Example 1: %v\n", result == expect)
}

func main() {

	ex1()
	ex2()
	ex3()
	ex4()
	ex5()

}

func restoreString(s string, indices []int) string {

	result := make([]rune, len(s))
	for i, c := range s {
		result[indices[i]] = c
	}

	return string(result)
}

/*
func restoreString(s string, indices []int) string {

    shuffle := make([]string, len(indices))

    for i, r := range s {
        newIndex := indices[i]
        shuffle[newIndex] = string(r)
    }

    var b strings.Builder
    for _, i := range shuffle {
        fmt.Fprintf(&b, "%s", string(i))
    }
    return b.String()
}
*/

/*
func restoreString(s string, indices []int) string {

    shuffle := make([]string, len(indices))

    for i, r := range s {
        shuffle[indices[i]] = string(r)
    }

    var b strings.Builder
    for _, i := range shuffle {
        fmt.Fprintf(&b, "%s", string(i))
    }
    return b.String()
}
*/

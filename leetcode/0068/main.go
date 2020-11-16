package main

import (
	"fmt"
	"strings"
)

func example1() {

	xwords := []string{"this", "is", "an", "example", "of", "text", "justification."}
	maxwidth := 16

	r := fullJustify(xwords, maxwidth)
	for _, v := range r {
		// fmt.Printf("key %v has string %v\n", k, v)
		fmt.Println(v)
	}
}

func example2() {

	xwords := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
		"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxwidth := 20

	r := fullJustify(xwords, maxwidth)
	for _, v := range r {
		// fmt.Printf("key %v has string %v\n", k, v)
		fmt.Println(v)
	}
}

func main() {
	example1()
	example2()
}

func fullJustify(words []string, maxWidth int) []string {

	var result []string
	var index int

	for index < len(words) {

		var wc int // word count
		var lc int // letter count
		var k int  // loop variable for index

		// line words are stored from words[index] to words[index + wc]
		for k = index; k < len(words); k++ {
			if lc+(wc)+len(words[k]) > maxWidth { // accomodate for at least 1 space between each word
				break
			}
			lc += len(words[k])
			wc++
		}

		// now just figure out spacing
		spacing := 1
		if wc > 1 {
			spacing = (maxWidth - lc) / (wc - 1)
		}

		var sep string
		for i := 0; i < spacing; i++ {
			sep += " "
		}

		temp := strings.Join(words[index:index+wc], sep)
		lpad := maxWidth - len(temp)
		var perm string
		for i := index; i < index+wc; i++ {
			if lpad > 0 {
				perm += words[i] + sep + " "
				lpad--
			} else {

				if i+1 == index+wc {
					perm += words[i]
				} else {
					perm += words[i] + sep
				}
			}
		}

		result = append(result, perm)

		index = k
	}

	return result
}

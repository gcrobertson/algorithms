package main

// Given a digit string, return all possible letter combinations that the number could represent.
// A mapping of digit to letters (just like on the telephone buttons) is given below.
//
// input: 	"23"
// output:	[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}

import (
	"fmt"
	"strconv"
)

// "2" ➜ "abc"
func digitMap(n string) string {

	var digitMap = map[int]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}

	i, _ := strconv.Atoi(string(n))

	return digitMap[i]
}

func bfs(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	var queue []string

	chars := digitMap(string(digits[0]))
	for _, v := range chars {
		queue = append(queue, string(v))
	}

	digits = digits[1:]

	for {

		if len(digits) == 0 {
			return queue
		}

		chars := digitMap(string(digits[0]))
		qsize := len(queue)

		for qsize > 0 {
			for i := 0; i < len(chars); i++ {

				// dequeue
				dq := queue[0]
				queue = queue[1:]

				for _, v := range chars {
					// enqueue
					queue = append(queue, dq+string(v))
				}
				qsize--
			}
		}

		digits = digits[1:]
	}
}

func main() {

	s := "234"

	r := bfs(s)
	for _, v := range r {
		fmt.Printf("%v\t", v)
	}
}

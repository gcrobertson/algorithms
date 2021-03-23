package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// Input: logs = []
	// Output: ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
	input := []string{
		"dig1 8 1 5 1",
		"let1 art can",
		"dig2 3 6",
		"let2 own kit dig",
		"let3 art zero",
	}

	expect := []string{
		"let1 art can",
		"let3 art zero",
		"let2 own kit dig",
		"dig1 8 1 5 1",
		"dig2 3 6",
	}

	output := reorderLogFiles(input)

	fmt.Printf("Does result : [%+v] match expectation [%+v]? [%v]\n", output, expect, reflect.DeepEqual(output, expect))
	// Input: logs = []
	// Output: ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
	input2 := []string{
		"a1 9 2 3 1",
		"g1 act car",
		"zo4 4 7",
		"ab1 off key dog",
		"a8 act zoo",
	}

	expect2 := []string{
		"g1 act car",
		"a8 act zoo",
		"ab1 off key dog",
		"a1 9 2 3 1",
		"zo4 4 7",
	}

	output2 := reorderLogFiles(input2)

	fmt.Printf("Does result : [%+v] match expectation [%+v]? [%v]\n", output2, expect2, reflect.DeepEqual(output2, expect2))

}

func reorderLogFiles(logs []string) []string {

	var llogs xLogs
	dlogs := []string{}

	for _, v := range logs {
		t1 := strings.Split(v, " ")
		if _, err := strconv.Atoi(t1[1]); err == nil {
			dlogs = append(dlogs, v)
		} else {
			llogs = append(llogs, v)
		}
	}

	// arrange the letter logs
	sort.Sort(llogs)

	// append the digital logs to the sorted letter logs
	// llogs = append(llogs, dlogs...)
	for _, v := range dlogs {
		llogs = append(llogs, v)
	}

	return llogs
}

type xLogs []string

func (s xLogs) Len() int {
	return len(s)
}
func (s xLogs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s xLogs) Less(i, j int) bool {

	t1 := strings.Split(s[i], " ")
	t2 := strings.Split(s[j], " ")

	c1 := strings.Join(t1[1:], "")
	c2 := strings.Join(t2[1:], "")

	return c1 < c2
}

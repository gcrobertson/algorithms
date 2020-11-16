package main

import (
	"fmt"
	"sort"
)

// Runtime: 80 ms, faster than 100.00% of Go online submissions for Boats to Save People.
// Memory Usage: 7.1 MB, less than 100.00% of Go online submissions for Boats to Save People.

func main() {

	// input := []int{5, 1, 4, 2}
	// limit := 6

	input := []int{3, 2, 2, 1}
	limit := 3

	result := numRescueBoats(input, limit)

	fmt.Printf("The answer I got is: %v\n", result)

}

func numRescueBoats(people []int, limit int) int {

	numberOfPeople := len(people)

	if numberOfPeople < 2 {
		return numberOfPeople
	}

	sort.Slice(people, func(x, y int) bool {
		return people[x] < people[y]
	})

	var left, boats int
	right := numberOfPeople - 1

	for left <= right {
		if people[left]+people[right] <= limit {
			left++
		}
		right--
		boats++
	}
	return boats
}

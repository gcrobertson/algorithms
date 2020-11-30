package main

import (
	"strconv"
	//"math"
)

func main() {

	//fmt.Println("Hello, playground")
	findMinDifference([]string{"23:59", "00:00", "00:16", "01:15", "02:46"})
}

func convert(input string) int {
	minutes, err := strconv.Atoi(input[0:2])
	minute, err := strconv.Atoi(input[3:5])
	if err != nil {
		return 0
	}
	// fmt.Printf("1: %v 2: %v\n", minutes, minute)
	return minutes*60 + minute
}

func insertion(convert int, sortedtime []int) []int { //O (log 2 n)

	if len(sortedtime) == 0 {
		sortedtime = append(sortedtime, convert)
		minimum = convert
		return sortedtime
	} else if len(sortedtime) == 1 {
		if sortedtime[0] < convert {
			sortedtime = append(sortedtime, convert)

		} else {
			sortedtime = []int{convert, sortedtime[0]}
		}
		minimum = sortedtime[1] - sortedtime[0]
		return sortedtime
	}

	// logarithmic insert below
	min := 0
	max := len(sortedtime) - 1
	for min < max {
		mid := max / 2
		if convert > sortedtime[mid] {
			min = mid
			if min == 1 {
				min++
			}
		} else if convert < sortedtime[mid] {
			max = mid
		}

		//fmt.Printf("min: %v\n max: %v\n", min, max)

		if max-min == 1 {

			sortedtime = append(sortedtime[:min+1], sortedtime[min:]...)
			sortedtime[min+1] = convert

			compareLow := min
			compareHigh := min + 2

			if sortedtime[compareHigh]-convert < minimum {
				minimum = sortedtime[compareHigh] - convert
			} else if convert-sortedtime[compareLow] < minimum {
				minimum = convert - sortedtime[compareLow]
			}

			return sortedtime
		}
	}

	return sortedtime
}

var minimum int

func findMinDifference(timePoints []string) int {

	timemap := make(map[int]bool, len(timePoints))
	var sortedtime []int

	for _, stringhours := range timePoints { // O(n)

		convert := convert(stringhours)

		if _, ok := timemap[convert]; ok {
			return 0
		}
		timemap[convert] = true
		sortedtime = insertion(convert, sortedtime)

	}

	return minimum
}

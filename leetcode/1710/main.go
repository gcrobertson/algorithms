package main

import "sort"

func main() {
}

func maximumUnits(boxTypes [][]int, truckSize int) int {

	// Need to sort the boxTypes by their value...
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	var result int

	for truckSize > 0 && len(boxTypes) > 0 {

		dq := boxTypes[0]
		boxTypes = boxTypes[1:]

		if truckSize-dq[0] > 0 {

			result += dq[0] * dq[1]
			truckSize -= dq[0]

		} else {
			result += truckSize * dq[1]
			return result
		}
	}

	return result
}

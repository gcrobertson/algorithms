package main

import "math"

func main() {

}

func partitionLabels(S string) []int {

	// map ["c"][end int]
	cmap := make(map[string]int, 26)

	// O(n)
	for key, val := range S {
		cmap[string(val)] = key
	}

	// cmap: map[a:8 b:5 c:7 d:14 e:15 f:11 g:13 h:19 i:22 j:23 k:20 l:21]

	var max int
	var anchor int
	var result []int

	// O(n)
	for index, letter := range S {
		max = int(math.Max(float64(max), float64(cmap[string(letter)])))
		// "ababcbaca"
		// max for index 0: 'a' := 0,8, max = 8
		// max for index 1: 'b' := 8,5, max = 8
		// max for index 2: 'a' := 8,8, max = 8
		// max for index 3: 'b' := 8,5, max = 8
		// max for index 4: 'c' := 8,7, max = 8
		// max for index 5: 'b' := 8,8, max = 8
		// max for index 6: 'a' := 8,5, max = 8
		// max for index 7: 'c' := 8,7, max = 8
		// max for index 8: 'a' := 8,8, max = 8

		// index 8 == max 8
		// result = []int{8 - 0 + 1} = 9
		// anchor = 8+1 = 9

		// "defegde"
		// max for index  9: 'd' := 8, 14, max = 14
		// max for index 10: 'e' := 14,15, max = 15
		// max for index 11: 'f' := 15,11  max = 15
		// max for index 12: 'e' := 15,15, max = 15
		// max for index 13: 'g' := 15,13, max = 15
		// max for index 14: 'd' := 15,14, max = 15
		// max for index 15: 'e' := 15,15, max = 15

		// index 15 == max 15
		// add 15 - 9 + 1
		// result = []int{9, 7}
		// anchor = 15 + 1 = 16

		// "hijhklij"
		// max for index 16: 'h' := 15,19, max = 19
		// max for index 17: 'i' := 19,22, max = 22
		// max for index 18: 'j' := 22,23, max = 23
		// max for index 19: 'h' := 23,19, max = 23
		// max for index 20: 'k' := 23,20, max = 23
		// max for index 21: 'l' := 23,21, max = 23
		// max for index 22: 'i' := 23,22, max = 23
		// max for index 23: 'j' := 23,23, max = 23

		// index 23 == max 23
		// add 23 - 16 + 1 = 8
		// result = []int{9,7,8}

		if index == max {
			result = append(result, index-anchor+1)
			anchor = index + 1
		}
	}

	// solution is O(n+n) = O(n)

	return result
}

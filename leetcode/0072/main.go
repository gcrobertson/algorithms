package main

import (
	"fmt"
	"math"
)

func main() {

	word1 := "Honda"
	word2 := "Hyundai"

	fmt.Printf("The minimum distance between %v and %v is %v\n", word1, word2, minDistance(word1, word2))
}

func minDistance(word1 string, word2 string) int {

	//      horse           hor
	//      one could compute: d[i][j]
	//      min(left, min(down, left down))
	//
	//      d[i-1][j]       // left
	//      d[i][j-1]       // down
	//      d[i-1][j-1]     // left down

	//              H   Y   U   N   D   A   I
	//          0   1   2   3   4   5   6   7
	//      H   1   0   1   2   3   4   5   6
	//      O   2   1   1   2   3   4   5   6
	//      N   3   2   2   2   2   3   4   5
	//      D   4   3   3   3   3   2   3   4
	//      A   5   4   4   4   4   3   2   3

	n := len(word1)
	m := len(word2)

	if n*m == 0 {
		return n + m
	}

	d := make([][]int, n+1)
	for i := 0; i < len(d); i++ {
		d[i] = make([]int, m+1)
	}

	// init boundaries
	for i := 0; i < n+1; i++ {
		d[i][0] = i
	}
	for j := 0; j < m+1; j++ {
		d[0][j] = j
	}

	for i := 1; i < n+1; i++ { // d[i+1] is the ith element

		for j := 1; j < m+1; j++ { // d[i+1][j+1] is the kth element

			left := d[i-1][j] + 1
			down := d[i][j-1] + 1
			diagonal := d[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				diagonal++
			}
			d[i][j] = int(math.Min(float64(left), math.Min(float64(down), float64(diagonal))))
		}
	}

	return d[n][m]

}

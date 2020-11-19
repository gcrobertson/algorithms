package main

import (
	"fmt"
)

func main() {
	image := [][]int{
		[]int{1, 1, 1},
		[]int{1, 1, 0},
		[]int{1, 0, 1},
	}
	sr := 1
	sc := 1

	newColor := 2

	result := floodFill(image, sr, sc, newColor)

	fmt.Printf("the result is %v\n", result)
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	originalColor := image[sr][sc]

	if originalColor == newColor {
		return image
	}

	floodFillColor(image, sr, sc, newColor, originalColor)

	return image
}

func floodFillColor(image [][]int, sr int, sc int, newColor int, originalColor int) {

	if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[0]) || image[sr][sc] != originalColor {
		return
	}

	image[sr][sc] = newColor

	floodFillColor(image, sr, sc-1, newColor, originalColor) // left
	floodFillColor(image, sr, sc+1, newColor, originalColor) // right
	floodFillColor(image, sr-1, sc, newColor, originalColor) // up
	floodFillColor(image, sr+1, sc, newColor, originalColor) // down
}

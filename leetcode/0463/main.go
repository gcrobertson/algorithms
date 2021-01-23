package main

import "fmt"

func main() {

	ex1()
	ex2()
	ex3()
}

func ex1() {
	input := [][]int{
		[]int{0, 1, 0, 0},
		[]int{1, 1, 1, 0},
		[]int{0, 1, 0, 0},
		[]int{1, 1, 0, 0},
	}
	expect := 16
	result := islandPerimeter(input)
	fmt.Printf("Does the result [%v] match the expecation? [%v]\n", result, expect == result)
}

func ex2() {
	input := [][]int{
		[]int{1},
	}
	expect := 4
	result := islandPerimeter(input)
	fmt.Printf("Does the result [%v] match the expecation? [%v]\n", result, expect == result)
}

func ex3() {
	input := [][]int{
		[]int{1, 0},
	}
	expect := 4
	result := islandPerimeter(input)
	fmt.Printf("Does the result [%v] match the expecation? [%v]\n", result, expect == result)
}

// Runtime: 64 ms, faster than 55.43% of Go online submissions for Island Perimeter.
// Memory Usage: 6.5 MB, less than 81.52% of Go online submissions for Island Perimeter.
func islandPerimeterSolution1(grid [][]int) int {

	var perimeter int

	for i := 0; i < len(grid); i++ { // rows

		for j := 0; j < len(grid[0]); j++ { // columns

			if grid[i][j] == 0 {
				continue
			}

			// now we know that we have land.
			perimeter += testPerimeter(grid, i-1, j) // left
			perimeter += testPerimeter(grid, i+1, j) // right
			perimeter += testPerimeter(grid, i, j-1) // bottom
			perimeter += testPerimeter(grid, i, j+1) // top
		}
	}

	return perimeter
}

// tests all 4 sides of the land ...
func testPerimeter(grid [][]int, i, j int) int {
	if i >= len(grid) || j >= len(grid[0]) || i < 0 || j < 0 || grid[i][j] == 0 { // grid boundaries will always be a perimeter as will water
		return 1
	}

	return 0
}

// Rather than checking 4 surrounding neighbors, we only need to check 2 neighbors (LEFT and UP) in this approach.
// Runtime: 64 ms, faster than 55.43% of Go online submissions for Island Perimeter.
// Memory Usage: 6.5 MB, less than 81.52% of Go online submissions for Island Perimeter.
// This solution is really no faster at all.. and makes sense because it is O(n*m) as well.
func islandPerimeterSolution2(grid [][]int) int {

	rows := len(grid)
	cols := len(grid[0])

	var perimeter int

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1 {

				perimeter += 4

				// UP
				if row > 0 && grid[row-1][col] == 1 {
					perimeter -= 2
				}

				// LEFT
				if col > 0 && grid[row][col-1] == 1 {
					perimeter -= 2
				}
			}
		}
	}
	return perimeter
}

// Adding an additional condition to the beginning can slightly speed up the runtime:
// Runtime: 60 ms, faster than 71.74% of Go online submissions for Island Perimeter.
// Memory Usage: 6.5 MB, less than 81.52% of Go online submissions for Island Perimeter.
func islandPerimeter(grid [][]int) int {

	rows := len(grid)
	cols := len(grid[0])

	if rows == 1 && cols == 1 {
		if grid[0][0] == 1 {
			return 4
		}
		if grid[0][0] == 0 {
			return 0
		}
	}

	var perimeter int

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1 {

				perimeter += 4

				// UP
				if row > 0 && grid[row-1][col] == 1 {
					perimeter -= 2
				}

				// LEFT
				if col > 0 && grid[row][col-1] == 1 {
					perimeter -= 2
				}
			}
		}
	}
	return perimeter
}

package main

import "fmt"

func main() {

	input := [][]int{
		[]int{1, 1, 0, 0, 0},
		[]int{1, 1, 0, 0, 0},
		[]int{0, 0, 0, 1, 1},
		[]int{0, 0, 0, 1, 1},
	}

	result := maxAreaOfIsland(input)

	fmt.Printf("The result is %v\n", result)
}

var maxArea int
var tempArea int

func maxAreaOfIsland(grid [][]int) int {

	maxArea = 0
	tempArea = 0

	if len(grid) == 0 {
		return 0
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] == 1 {
				// grid[i][j] = 0  // set island to water
				tempArea = 0
				maxAreaOfIslandDFS(i, j, grid)
				if tempArea > maxArea {
					maxArea = tempArea
				}
			}
		}
	}

	return maxArea
}

func maxAreaOfIslandDFS(x int, y int, grid [][]int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return
	}

	if grid[x][y] != 1 {
		return
	}

	grid[x][y] = 0 // set island to water
	tempArea++

	// fmt.Printf("tempArea: %v for grid[%v][%v]\n", tempArea, x, y)

	// go down
	maxAreaOfIslandDFS(x, y-1, grid)

	// go up
	maxAreaOfIslandDFS(x, y+1, grid)

	// go left
	maxAreaOfIslandDFS(x-1, y, grid)

	// go right
	maxAreaOfIslandDFS(x+1, y, grid)
}

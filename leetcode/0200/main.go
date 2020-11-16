package main

import "fmt"

func main() {

	grid := [][]byte{
		// []byte{'1', '0', '1'},
		// []byte{'1', '1', '0'},
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}

	r := numIslands(grid)

	fmt.Printf("The number of islands: %d\n", r)
}

func numIslands(grid [][]byte) int {

	if len(grid) == 0 {
		return 0
	}

	var zeros int

	uf := initUF(len(grid) * len(grid[0]))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '0' {
				zeros++
				continue
			}
			if j > 0 && grid[i][j] == '1' && grid[i][j-1] == '1' {
				uf.Union()
			}
		}
	}

	// https://leetcode.com/problems/number-of-islands/discuss/542463/Golang-100-weighted-union-find

}

type unionFind struct {
	id    []int
	size  []int
	Count int
}

func initUF(size int) *unionFind {
	uf := unionFind{
		id:    make([]int, size),
		size:  make([]int, size),
		Count: size,
	}
	for i := 0; i < size; i++ {
		uf.id[i] = i
		uf.size[i] = 1
	}
	return &uf
}

/** DFS Submission below:
func dfs(grid *[][]byte, row, col int) {
	if row < 0 || col < 0 || row >= len(*grid) || col >= len((*grid)[0]) || (*grid)[row][col] == '0' {
		return
	}
	(*grid)[row][col] = '0'
	dfs(grid, row+1, col)
	dfs(grid, row-1, col)
	dfs(grid, row, col+1)
	dfs(grid, row, col-1)

}

func numIslands(grid [][]byte) int {

	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	var islands int

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				islands++
				dfs(&grid, i, j)
			}
		}
	}

	return islands
}
*/

/** BFS Submission below:

var directions = [][]int{
	{1, 0},  // right
	{-1, 0}, // left
	{0, 1},  // up
	{0, -1}, // down
}

// 2D to 1D: row number * number of columns + coordinate value
// 1D to 2D:
//	Row: 	coordinate value / number of columns
//  Column: coordinate value % number of columns
func bfs(grid [][]byte, rows, cols, i, j int) {

	queue := []int{}
	queue = append(queue, i*cols+j) // add to queue as 1D
	grid[i][j] = '0'                // convert into water

	for len(queue) > 0 {
		index := queue[0]
		queue = queue[1:]

		row := index / cols
		col := index % cols

		for _, coordinates := range directions {
			r := row + coordinates[0]
			c := col + coordinates[1]
			if r > -1 && c > -1 && r < rows && c < cols && grid[r][c] == '1' {

				grid[r][c] = '0'

				oneD := r*cols + c

				queue = append(queue, oneD)
			}
		}
	}
}

// convert back from int to byte for leetcode submission
func numIslands(grid [][]byte) int {

	if len(grid) == 0 {
		return 0
	}

	var islands int

	rows := len(grid)
	cols := len(grid[0])

	// loop over every element in 2D array
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			if grid[i][j] == '1' { // we only care about land
				islands++
				// start performing BFS
				bfs(grid, rows, cols, i, j)
			}
		}
	}

	return islands
}

**/

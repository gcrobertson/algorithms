package main

import "fmt"

func main() {

	input := [][]int{
		[]int{0, 1, 0},
		[]int{0, 0, 1},
		[]int{1, 1, 1},
		[]int{0, 0, 0},
	}

	gameOfLife(input)

	fmt.Printf("Result: %+v\n", input)
}

func gameOfLife(board [][]int) {

	directions := [][]int{
		// up
		[]int{-1, 0},
		// up left
		[]int{-1, -1},
		// up right
		[]int{-1, 1},
		// left
		[]int{0, -1},
		// right
		[]int{0, 1},
		// down left
		[]int{1, -1},
		// down
		[]int{1, 0},
		// down right
		[]int{1, 1},
	}

	// no DFS
	// can solve in O(n)

	newBoard := make([][]int, len(board))

	for i := 0; i < len(newBoard); i++ {
		newBoard[i] = make([]int, len(board[0]))
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			// we are iterating through all board[row][col] at this point O(n)

			coordinateState := board[row][col]
			counter := 0

			for _, direction := range directions {

				newRow := row + direction[0]
				newCol := col + direction[1]

				if newRow >= 0 && newCol >= 0 && newRow < len(board) && newCol < len(board[0]) {

					if board[newRow][newCol] == 1 {
						counter++
					}
				}
			}

			//fmt.Printf("the coordinate [%v,%v] has a counter of %v\n", row, col, counter)

			// Apply the rules

			// 1. Any live cell with fewer than two live neighbors dies
			if coordinateState == 1 && counter < 2 {
				newBoard[row][col] = 0
			} else if coordinateState == 1 && counter > 1 && counter < 4 {
				// 2. Any live cell with two or three live neighbors lives on to the next generation
				newBoard[row][col] = 1
			} else if counter > 3 {
				// 3. Any live cell with more than three live neighbors dies
				newBoard[row][col] = 0
			} else if coordinateState == 0 && counter == 3 {
				newBoard[row][col] = 1
			}
		}
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			board[row][col] = newBoard[row][col]
		}
	}

	// fmt.Println("newBoard", newBoard)
}

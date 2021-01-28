package main

import "fmt"

func main() {

	ex1()
	ex2()
	ex3()
}

func ex1() {

	input := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}

	expect := 3

	result := numRookCaptures(input)

	fmt.Printf("result [%v] matches expectation? [%v]\n", result, result == expect)
}

func ex2() {

	input := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'B', 'R', 'B', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}

	expect := 0

	result := numRookCaptures(input)

	fmt.Printf("result [%v] matches expectation? [%v]\n", result, result == expect)
}

func ex3() {

	input := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'}, // [".",".",".",".",".",".",".","."]
		{'.', '.', '.', 'p', '.', '.', '.', '.'}, // [".",".",".","p",".",".",".","."]
		{'.', 'p', 'p', 'p', '.', '.', '.', '.'}, // [".",".",".","p",".",".",".","."]
		{'p', 'p', '.', 'R', '.', 'p', 'B', '.'}, // ["p","p",".","R",".","p","B","."]
		{'.', '.', '.', '.', '.', '.', '.', '.'}, // [".",".",".",".",".",".",".","."]
		{'.', '.', '.', 'B', '.', '.', '.', '.'}, // [".",".",".","B",".",".",".","."]
		{'.', '.', '.', 'p', '.', '.', '.', '.'}, // [".",".",".","p",".",".",".","."]
		{'.', '.', '.', '.', '.', '.', '.', '.'}, // [".",".",".",".",".",".",".","."]
	}

	expect := 3

	result := numRookCaptures(input)

	fmt.Printf("result [%v] matches expectation? [%v]\n", result, result == expect)
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Available Captures for Rook.
// Memory Usage: 1.9 MB, less than 8.70% of Go online submissions for Available Captures for Rook.
func numRookCaptures(board [][]byte) int {

	var rooky int
	var rookx int
	var result int

	// O(n*m)
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if board[y][x] == 'R' {
				rooky = y
				rookx = x
			}
		}
	}

	// check vertically up
	// O(n)
	for y := rooky; y > -1; y-- {
		if board[y][rookx] == 'B' {
			break
		} else if board[y][rookx] == 'p' {
			result++
			break
		}
	}

	// check vertically down
	// O(n)
	for y := rooky; y < len(board[0]); y++ {
		if board[y][rookx] == 'B' {
			break
		} else if board[y][rookx] == 'p' {
			result++
			break
		}
	}

	// check horizontally left
	// O(m)
	for x := rookx; x > -1; x-- {
		if board[rooky][x] == 'B' {
			break
		} else if board[rooky][x] == 'p' {
			result++
			break
		}
	}

	// check horizontally right
	// O(m)
	for x := rookx; x < len(board); x++ {
		if board[rooky][x] == 'B' {
			break
		} else if board[rooky][x] == 'p' {
			result++
			break
		}
	}

	return result
}

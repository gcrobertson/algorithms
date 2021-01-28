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
	var v1 int
	for y := 0; y < rooky; y++ {
		if board[y][rookx] == 'B' {
			v1 = 0
		} else if board[y][rookx] == 'p' && v1 == 0 {
			v1++
		}
	}

	// check vertically down
	// O(n)
	var v2 int
	for y := len(board[0]) - 1; y > rooky; y-- {
		if board[y][rookx] == 'B' {
			v2 = 0
		} else if board[y][rookx] == 'p' && v2 == 0 {
			v2++
		}
	}

	// check horizontally left
	// O(m)
	var h1 int
	for x := 0; x < rookx; x++ {
		if board[rooky][x] == 'B' {
			h1 = 0
		} else if board[rooky][x] == 'p' && h1 == 0 {
			h1++
		}
	}

	// check horizontally right
	// O(m)
	var h2 int
	for x := len(board) - 1; x > rookx; x-- {
		if board[rooky][x] == 'B' {
			h2 = 0
		} else if board[rooky][x] == 'p' && h2 == 0 {
			h2++
		}
	}

	return v1 + v2 + h1 + h2
}

package main

import "fmt"

func main() {

	input := "GGLLGG"

	output := isRobotBounded(input)

	fmt.Printf("Is the robot bounded? %v\n", output)

}

func isRobotBounded(instructions string) bool {

	// north = 0, east = 1, south = 2, west = 3
	directions := [][]int{
		[]int{0, 1},
		[]int{1, 0},
		[]int{0, -1},
		[]int{-1, 0},
	}

	var x, y int
	var idx int // north 0

	for _, v := range instructions {
		switch v {
		case 'G': // straight
			x += directions[idx][0]
			y += directions[idx][1]
		case 'L': // rotate left
			idx = (idx + 3) % 4
		case 'R': // rotate right
			idx = (idx + 1) % 4
		}
	}

	if (x == 0 && y == 0) || idx != 0 {
		return true
	}
	return false
}

package main

func main() {

}

func checkStraightLine(coordinates [][]int) bool {
	coordinate2 := coordinates[1]
	coordinate1 := coordinates[0]
	y1 := coordinate2[1]
	y0 := coordinate1[1]
	x1 := coordinate2[0]
	x0 := coordinate1[0]
	// slope := (y2 - y1) / (x2 - x1)
	if len(coordinates) == 2 {
		return true
	}
	for i := 2; i < len(coordinates); i++ {
		y2 := coordinates[i][1]
		x2 := coordinates[i][0]

		if ((y2 - y1) * (x2 - x0)) != ((x2 - x1) * (y2 - y0)) {
			return false
		}
	}

	return true
}

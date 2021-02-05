package main

func main() {

}

func numJewelsInStones(jewels string, stones string) int {

	if len(stones) == 0 {
		return 0
	}

	j := make(map[rune]bool)

	for _, jewel := range jewels {
		j[jewel] = true
	}

	var result int
	for _, stone := range stones {
		if _, ok := j[stone]; ok {
			result++
		}
	}

	return result
}

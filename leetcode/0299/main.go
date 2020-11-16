package main

import (
	"fmt"
	"strconv"
)

func main() {

	// input := "1807"
	// guess := "7810"

	// input := "1123"
	// guess := "0111"

	// input := "1122"
	// guess := "1222"
	// hint1 := "3A0B"

	input := "00112233445566778899"
	guess := "16872590340158679432"
	hint1 := "3A17B"

	hint := getHint(input, guess)

	fmt.Printf("%v %v\n", hint, hint == hint1)
}

// O(n+n)
func getHint(secret string, guess string) string {
	// bulls := match value & key
	// cows := match value but not key
	secretMap := make(map[string]int)
	for i := 0; i < len(secret); i++ { // O(n)
		secretMap[string(secret[i])]++
	}
	var bulls, cows int
	// assume same lengths secret & guess
	for i := 0; i < len(secret); i++ { // O(n)
		g := string(guess[i])
		s := string(secret[i])
		if g == s {
			bulls++
			secretMap[g]--
		}
	}

	for i := 0; i < len(secret); i++ {
		g := string(guess[i])
		s := string(secret[i])
		if g != s {
			if count, ok := secretMap[g]; ok {
				if count > 0 {
					cows++
					secretMap[g]--
				}
			}
		}

	}

	// return FormatInt(int64(bulls), 10) + "A" + FormatInt(int64(cows), 10) + "B"
	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}

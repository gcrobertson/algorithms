package main

import "math"

func main() {

}

func myAtoi(str string) int {

	var i int // starts at 0
	flag := 1

	// remove all leading whitespace
	for i < len(str) {
		if str[i] != ' ' {
			break
		}
		i++
	}

	// handle empty or whitespace strings
	if i == len(str) || len(str) == 0 {
		return 0
	}

	// handle mathematical symbol [+/-]
	if str[i] == '-' {
		flag = -1
		i++
	} else if str[i] == '+' {
		i++
	}

	var val int

	// traverse over all valid characters in ASCII range of 48-57 [0-9]
	for ; i < len(str) && '0' <= str[i] && str[i] <= '9'; i++ {
		val = val*10 + int(str[i]-'0') // A much cleaner version of what I tried to show you yesterday
		if flag*val > math.MaxInt32 {
			return math.MaxInt32
		} else if flag*val < math.MinInt32 {
			return math.MinInt32
		}
	}
	return flag * val
}

package main

func main() {

}

func romanToInt(s string) int {

	n := len(s)
	if n < 1 {
		return 0
	}

	table := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var val, i int

	for i < n {
		if i+1 < n && table[rune(s[i])] < table[rune(s[i+1])] {
			val += table[rune(s[i+1])] - table[rune(s[i])]
			i += 2
		} else {
			val += table[rune(s[i])]
			i++
		}
	}
	return val
}

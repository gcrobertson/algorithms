package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	result := addStrings("11", "123")
	expect := "134"
	fmt.Printf("%v\t%v\t%v\n", expect, result, result == expect)

	result = addStrings("456", "77")
	expect = "533"
	fmt.Printf("%v\t%v\t%v\n", expect, result, result == expect)
}

// func addStrings(num1 string, num2 string) string {

// 	p1, p2 := len(num1)-1, len(num2)-1
// 	var result string //s.Builder
// 	var carry int

// 	for p1 >= 0 || p2 >= 0 || carry != 0 {
// 		sum := carry
// 		if p1 >= 0 {
// 			sum += int(num1[p1] - '0')
// 			p1--
// 		}
// 		if p2 >= 0 {
// 			sum += int(num2[p2] - '0')
// 			p2--
// 		}
// 		carry = sum / 10
// 		// fmt.Fprintf(&result, "%v", sum)
// 		result = strconv.Itoa(sum%10) + result
// 	}

// 	return result // reverse(result.String())
// }

func reverse(s string) string {
	var p1 int
	p2 := len(s) - 1
	rns := []rune(s)
	for p1 < p2 {
		rns[p1], rns[p2] = rns[p2], rns[p1]
		p1++
		p2--
	}
	return string(rns)
}

// func addStrings2(num1 string, num2 string) string {
// 	x, y := len(num1)-1, len(num2)-1
// 	res, carry := "", 0
// 	for x >= 0 || y >= 0 || carry != 0 {
// 		sum := carry
// 		if x >= 0 {
// 			sum += int(num1[x] - '0')
// 			x--
// 		}
// 		if y >= 0 {
// 			sum += int(num2[y] - '0')
// 			y--
// 		}
// 		carry = sum / 10
// 		res = strconv.Itoa(sum%10) + res
// 	}
// 	return res
// }

func addStrings(num1 string, num2 string) string {
	var carryOver int
	var sb strings.Builder

	p1 := len(num1) - 1
	p2 := len(num2) - 1

	for p1 >= 0 || p2 >= 0 || carryOver != 0 {
		var d1, d2 int
		if p1 >= 0 {
			d2 = int(num1[p1] - '0')
			p1--
		}

		if p2 >= 0 {
			d1 = int(num2[p2] - '0')
			p2--
		}

		digit := d1 + d2 + carryOver
		if digit >= 10 {
			sb.WriteString(strconv.Itoa(digit - 10))
			carryOver = 1
		} else {
			sb.WriteString(strconv.Itoa(digit))
			carryOver = 0
		}
	}
	return reverse(sb.String())
}

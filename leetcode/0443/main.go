package main

import "fmt"

func main() {

	chars := []byte{97, 97, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 99, 99, 99} //aabbccc			//a2b2c3

	result := compress(chars)

	expect := 6

	fmt.Printf("Does the result [%v] match the expectation [%v]? [%v]\n", result, expect, result == expect)

	chars = chars[0:result]

	fmt.Println(string(chars))

}

func compress(chars []byte) int {

	scan := 0
	write := 0

	for scan < len(chars) {
		count := 0
		chars[write] = chars[scan]
		for scan < len(chars) && chars[write] == chars[scan] {
			count++
			scan++
		}

		// breaking when the characters dont match... very clever
		if count > 1 {

			// this actually handles when the count is > 9 by storing the value as a string and then iterating over it below
			tmp := fmt.Sprintf("%d", count)
			for _, c := range []byte(tmp) {
				write++
				chars[write] = c
			}
		}
		write++
	}

	return write
}

/*
func main() {
	input := []byte{97, 97, 98, 98, 99, 99, 99} // "a", "2", "b", "2", "c", "3"	returning 6
	result := compress(input)
	expect := 6
	fmt.Printf("Does the result [%v] match the expectation [%v]? [%v]\n", result, expect, result == expect)
	fmt.Printf("value of input %v\n", string(input))
}

func compress(chars []byte) int {

	if len(chars) == 1 {
		return 1
	}

	result := []byte{}
	var total int

	for i := 1; i < len(chars); i++ {

		if total == 0 {
			result = append(result, chars[i]) // Append the new character
			total++
		}

		if chars[i] == chars[i-1] {
			// Same character as previous
			total++
		} else {

			// Different character as previous
			if total > 1 {
				t := strconv.Itoa(total)
				result = append(result, []byte(t)[0])
			}
			total = 0
		}

		if i+1 == len(chars) {
			t := strconv.Itoa(total)
			result = append(result, []byte(t)[0])
		}
	}

	//chars = make([]byte, len(result))

	for i := 0; i < len(result); i++ {
		chars[i] = result[i]
	}

	chars = chars[0:len(result)]

	fmt.Printf("value of chars %v\n", string(chars))

	return len(result)
}
*/

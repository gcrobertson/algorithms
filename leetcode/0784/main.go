package main

/*
func main() {

	input := "a1b2"

	output := []string{"a1b2", "a1B2", "A1b2", "A1B2"}

	result := lettercasePermutation(input)

	fmt.Printf("%v should match %v\n", output, result)

}

func lettercasePermutation(input string) []string {

	var characters []int
	fmt.Println(characters)

	for k, v := range input {
		if unicode.IsNumber(v) {
			continue
		}
		characters = append(characters, k)
	}

	if len(characters) == 0 {
		return []string{input}
	}

	result := []string{input}

	for i := 0; i < len(characters); i++ {

		fmt.Printf("%T %v", input[characters[i]], input[characters[i]])

		if unicode.IsLower(rune(input[characters[i]])) {

			fmt.Printf("the key %v is definitely lowercase: %v\n", i, input[i])

			s := input
			s[characters[i]] =

		}

		// else {
		// 	fmt.Printf("the key %v is definitely uppercase: %v\n", i, input[i])
		// }
	}

	return result
}
*/
// fmt.Printf("i know the characters in the passed string are here: %+v\n", characters)
// fmt.Println(string(input[1]))

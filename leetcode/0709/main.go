package main

import (
	"fmt"
	"strings"
)

func main() {

}

func toLowerCase(str string) string {

	var result string

	for _, s := range str {

		if s > 64 && s < 91 {
			s += 32
		}

		result += string(s)
	}

	return result
}

// Using a strings.Builder
func toLowerCaseSB(str string) string {

	var result strings.Builder

	for _, s := range str {

		if s > 64 && s < 91 {
			s += 32
		}

		fmt.Fprintf(&result, "%s", string(s))
	}

	return result.String()
}

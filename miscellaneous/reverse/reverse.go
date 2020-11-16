package main

import "fmt"

func main() {

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(s1)
	reverse(s1)
	fmt.Println(s1)

	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(s2)
	quadReverse(s2)
	fmt.Println(s2)
}

// Linear time complexity: Time complexity Θ(n)
func reverse(s []int) {
	var i int
	var j = len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

// Quadratic time complexity: Time complexity Θ(n2)
func quadReverse(s []int) {
	for i := 1; i < len(s); i++ {
		x := s[i]
		for j := i; j >= 1; j-- {
			s[j] = s[j-1]
		}
		s[0] = x
	}
}

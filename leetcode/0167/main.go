package main

func main() {

}

func twoSum(numbers []int, target int) []int {

	h := make(map[int]int) // value | index

	for i := 0; i < len(numbers); i++ {

		n := target - numbers[i]

		if index, ok := h[n]; ok {
			return []int{index + 1, i + 1}
		}

		h[numbers[i]] = i
	}

	return []int{-1, -1}
}

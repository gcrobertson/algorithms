package main

func main() {

}

func countSmaller(nums []int) []int {
	var i int
	result := make([]int, len(nums))
	for i < len(nums) {
		j := len(nums) - 1
		for j > i {
			if nums[j] < nums[i] {
				result[i]++
			}
			j--
		}
		i++
	}
	return result
}

package main

func main() {

}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Distribute Candies to People.
// Memory Usage: 2 MB, less than 27.27% of Go online submissions for Distribute Candies to People.
func distributeCandies(candies int, numberOfpeople int) []int {

	result := make([]int, numberOfpeople)

	index := 0
	give := 1

	for candies > 0 {

		if index == numberOfpeople {
			index = 0
		}

		if candies < give {
			give = candies
		}

		result[index] += give
		candies -= give

		give++
		index++
	}

	return result
}

/*
func distributeCandies(candies int, num_people int) []int {
	var res []int = make([]int,num_people,num_people)
	var cnt int = 1
	for candies > cnt{
		res[(cnt - 1) % num_people] += cnt
		candies -= cnt
		cnt++
	}
	res[(cnt - 1) % num_people] += candies
	return res
}
*/

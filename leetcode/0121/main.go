package main

import "math"

func main() {

}

func maxProfit(prices []int) int {
	var min, profit int
	for k, v := range prices {
		if k == 0 {
			min = v
			profit = 0
		} else {

			min = int(math.Min(float64(v), float64(min)))
			profit = int(math.Max(float64(v)-float64(min), float64(profit)))

			/*
			   t := v - min
			   if t > profit {
			       profit = t
			   }
			   if v < min {
			       min = v
			   }
			*/
		}
	}
	return profit
}

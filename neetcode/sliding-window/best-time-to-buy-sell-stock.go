package main

import (
	"fmt"
	"math"
)

func maxProfit(list []int) int {
	buyDay, sellDay := 0, 1
	max_profit := 0

	for sellDay < len(list) {
		diff := list[sellDay] - list[buyDay]
		if diff < 0 {
			buyDay = sellDay
		} else {

			max_profit = int(math.Max(float64(diff), float64(max_profit)))
		}
		sellDay += 1
	}

	return max_profit
}

func main() {
	fmt.Printf("maxProfit: %v\n", maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

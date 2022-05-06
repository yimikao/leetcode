package main

import "fmt"

func twoSum(arr []int, target int) (result []int) {

	hash_set := make(map[int]int)

	for indexCurrNum, num := range arr {
		diff := target - num
		if indexOfDiff, exists := hash_set[diff]; exists {
			return append(result, indexOfDiff, indexCurrNum)
		}
		hash_set[num] = indexCurrNum
	}
	return
}

func main() {

	target := 4

	arr := []int{2, 1, 5, 3}

	fmt.Println(twoSum(arr, target))

}

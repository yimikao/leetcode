package main

import "fmt"

func two_sum2(arr []int, target int) (result []int) {

	leftPivot, rightPivot := 0, len(arr)-1

	for leftPivot < rightPivot {

		currSum := arr[leftPivot] + arr[rightPivot]

		if currSum < target {
			leftPivot++
		} else if currSum > target {
			rightPivot--
		} else {
			return append(result, leftPivot+1, rightPivot+1)
		}
	}
	return
}

func main() {

	fmt.Println(two_sum2([]int{1, 2, 4, 7, 8, 9, 15}, 9))

}

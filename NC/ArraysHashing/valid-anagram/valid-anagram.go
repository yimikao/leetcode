package main

import "fmt"

/***

 */

func main() {
	fmt.Println(IsValidAnagram("lily", "lele"))
}

func IsValidAnagram(s, t string) bool {

	if len(s) != len(t) {
		return false
	}

	countS := map[string]int{}
	countT := map[string]int{}

	for i := range s {
		countS[string(s[i])]++
		countT[string(t[i])]++
	}

	for i := range countS {
		if countS[i] != countT[i] {
			return false
		}
	}
	return true
}

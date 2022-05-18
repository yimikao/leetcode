package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, c := range s {
		if ch, ok := closeToOpen[c]; ok {
			// that is c is a closing bracket bracket
			if len(stack) > 0 && ch == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		} else {
			stack = append(stack, c)
		}
	}
	return !(len(stack) != 0)
}

func main() {

	fmt.Printf("isValid(): %v\n", isValid("()["))
	fmt.Printf("isValid(): %v\n", isValid("{{()}}"))
}

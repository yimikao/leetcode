package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isAlphaNum(c byte) bool {
	// return c < '0'
	if unicode.IsSpace(rune(c)) {
		return false
	}
	return unicode.IsLetter(rune(c)) || unicode.IsNumber(rune(c))

}

func isPalindrome(s string) bool {
	// _ = strings.ToLower(*s)

	leftPivot, rightPivot := 0, len(s)-1

	for leftPivot < rightPivot {

		leftChar := s[leftPivot]
		rightChar := s[rightPivot]

		if !isAlphaNum(leftChar) {
			leftPivot++
		}
		if !isAlphaNum(rightChar) {
			rightPivot--

		}
		if !strings.EqualFold(strings.ToLower(string(s[leftPivot])),
			strings.ToLower(string(s[rightPivot]))) {

			return false
		}
		leftPivot++
		rightPivot--
	}
	return true

}

func main() {

	fmt.Println(isPalindrome("ma  dam"))

}

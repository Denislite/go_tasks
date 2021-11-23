package palindrome

import (
	"strings"
)

func palindromeCheck(input string) bool {
	input = strings.ToLower(input)
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

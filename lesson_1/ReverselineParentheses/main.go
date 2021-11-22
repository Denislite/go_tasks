package main

import (
	"fmt"
)

func reverseLine(input string) (result string) {
	for _, char := range input {
		result = string(char) + result
	}
	return result[1:]
}

func reverseParantheses(input string) string {
	lastIndex := 0
	for index := 0; index < len(input); index++ {
		if input[index] == ')' {
			input = input[:lastIndex] + reverseLine(input[lastIndex+1:index+1]) + input[index+1:]
			lastIndex = 0
			index = 0
		}
		if input[index] == '(' {
			lastIndex = index
		}
	}
	return input
}

func main() {
	test := "foo(bar(baz))blim"
	fmt.Println(reverseParantheses(test))
}

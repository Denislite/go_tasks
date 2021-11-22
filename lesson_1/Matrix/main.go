package main

import (
	"fmt"
)

type Matrix [][]int

func matrixElementSum(input Matrix) int {
	sum, rows, cols := 0, len(input)-1, len(input[0])-1
	for i := 0; i <= cols; i++ {
		for j := 0; j <= rows; j++ {
			if input[j][i] == 0 {
				break
			} else {
				sum += input[j][i]
			}
		}
	}
	return sum
}

func main() {
	arr := Matrix{{1, 3, 2, 0}, {0, 5, 0, 9}, {1, 2, 3, 4}}
	fmt.Println(matrixElementSum(arr))
}

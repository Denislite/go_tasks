package matrix

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

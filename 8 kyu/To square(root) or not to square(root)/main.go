package kata

import "math"

// SquareOrSquareRoot https://www.codewars.com/kata/to-square-root-or-not-to-square-root
func SquareOrSquareRoot(arr []int) []int {
	result := make([]int, len(arr))

	for i, number := range arr {
		squareRoot := math.Sqrt(float64(number))

		if math.Trunc(squareRoot) == squareRoot {
			result[i] = int(squareRoot)
		} else {
			result[i] = number * number
		}
	}

	return result
}

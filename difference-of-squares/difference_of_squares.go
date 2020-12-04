package diffsquares

//SquareOfSum returns the square of the sum
func SquareOfSum(n int) int {
	return (n * (n + 1) / 2) * (n * (n + 1) / 2)
}

//SumOfSquares returns the sum of all the squares up to n
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

//Difference calculates the difference between squaresofSum and sumofSquares
func Difference(n int) int {
	sumofSquares := SumOfSquares(n)
	squaresofSum := SquareOfSum(n)
	return squaresofSum - sumofSquares
}
